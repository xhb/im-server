package db

import (
	"Open_IM/pkg/common/config"
	log2 "Open_IM/pkg/common/log"
	pbChat "Open_IM/pkg/proto/chat"
	pbRtc "Open_IM/pkg/proto/rtc"
	pbCommon "Open_IM/pkg/proto/sdk_ws"
	"Open_IM/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"time"
)

//func  (d *  DataBases)pubMessage(channel, msg string) {
//   d.rdb.Publish(context.Background(),channel,msg)
//}
//func  (d *  DataBases)pubMessage(channel, msg string) {
//	d.rdb.Publish(context.Background(),channel,msg)
//}

func (d *DataBases) NewGetMessageListBySeq(userID string, seqList []uint32, operationID string) (seqMsg []*pbCommon.MsgData, failedSeqList []uint32, errResult error) {
	for _, v := range seqList {
		//MESSAGE_CACHE:169.254.225.224_reliability1653387820_0_1
		key := messageCache + userID + "_" + strconv.Itoa(int(v))

		result, err := d.rdb.HGetAll(context.Background(), key).Result()
		if err != nil {
			errResult = err
			failedSeqList = append(failedSeqList, v)
			log2.NewWarn(operationID, "redis get message error:", err.Error(), v)
		} else {
			msg, err := Map2Pb(result)
			//msg := pbCommon.MsgData{}
			//err = jsonpb.UnmarshalString(result, &msg)
			if err != nil {
				errResult = err
				failedSeqList = append(failedSeqList, v)
				log2.NewWarn(operationID, "Unmarshal err", result, err.Error())
			} else {
				log2.NewDebug(operationID, "redis get msg is ", msg.String())
				seqMsg = append(seqMsg, msg)
			}

		}
	}
	return seqMsg, failedSeqList, errResult
}
func Map2Pb(m map[string]string) (*pbCommon.MsgData, error) {
	var data pbCommon.MsgData
	err := mapstructure.Decode(m, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (d *DataBases) NewSetMessageToCache(msgList []*pbChat.MsgDataToMQ, uid string, operationID string) error {
	ctx := context.Background()
	var failedList []pbChat.MsgDataToMQ
	for _, msg := range msgList {
		key := messageCache + uid + "_" + strconv.Itoa(int(msg.MsgData.Seq))
		s, err := utils.Pb2Map(msg.MsgData)
		if err != nil {
			log2.NewWarn(operationID, utils.GetSelfFuncName(), "Pb2Map failed", msg.MsgData.String(), uid, err.Error())
			continue
		}
		log2.NewDebug(operationID, "convert map is ", s)
		fmt.Println("ts", s)
		err = d.rdb.HMSet(context.Background(), key, s).Err()
		//err = d.rdb.HMSet(context.Background(), "12", map[string]interface{}{"1": 2, "343": false}).Err()
		if err != nil {
			return err
			log2.NewWarn(operationID, utils.GetSelfFuncName(), "redis failed", "args:", key, *msg, uid, s, err.Error())
			failedList = append(failedList, *msg)
		}
		d.rdb.Expire(ctx, key, time.Second*time.Duration(config.Config.MsgCacheTimeout))
	}
	if len(failedList) != 0 {
		return errors.New(fmt.Sprintf("set msg to cache failed, failed lists: %q,%s", failedList, operationID))
	}
	return nil
}

func (d *DataBases) CleanUpOneUserAllMsgFromRedis(userID string, operationID string) error {
	ctx := context.Background()
	key := messageCache + userID + "_" + "*"
	vals, err := d.rdb.Keys(ctx, key).Result()
	log2.Debug(operationID, "vals: ", vals)
	if err == redis.ErrNil {
		return nil
	}
	if err != nil {
		return utils.Wrap(err, "")
	}
	if err = d.rdb.Del(ctx, vals...).Err(); err != nil {
		return utils.Wrap(err, "")
	}
	return nil
}

func (d *DataBases) NewCacheSignalInfo(msg *pbCommon.MsgData) error {
	keyList := SignalListCache + msg.RecvID
	timeout, err := strconv.Atoi(config.Config.Rtc.SignalTimeout)
	if err != nil {
		return err
	}
	err = d.rdb.LPush(context.Background(), keyList, msg.ClientMsgID).Err()
	if err != nil {
		return err
	}
	err = d.rdb.Expire(context.Background(), keyList, time.Duration(timeout)*time.Second).Err()
	if err != nil {
		return err
	}
	key := SignalCache + msg.ClientMsgID
	err = d.rdb.Set(context.Background(), key, msg.Content, time.Duration(timeout)*time.Second).Err()
	if err != nil {
		return err
	}
	err = d.rdb.Expire(context.Background(), key, time.Duration(timeout)*time.Second).Err()
	if err != nil {
		return err
	}
	return err
}

func (d *DataBases) GetSignalInfoFromCacheByClientMsgID(clientMsgID string) (invitationInfo *pbRtc.SignalInviteReq, err error) {
	key := SignalCache + clientMsgID
	bytes, err := d.rdb.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}
	req := &pbRtc.SignalReq{}
	if err = proto.Unmarshal(bytes, req); err != nil {
		return nil, err
	}
	req2 := req.Payload.(*pbRtc.SignalReq_Invite)
	invitationInfo = req2.Invite
	return invitationInfo, err
}

func (d *DataBases) GetAvailableSignalInvitationInfo(userID string) (invitationInfo *pbRtc.SignalInviteReq, err error) {
	keyList := SignalListCache + userID
	result := d.rdb.RPop(context.Background(), keyList)
	if err = result.Err(); err != nil {
		return nil, err
	}
	invitationInfo, err = d.GetSignalInfoFromCacheByClientMsgID(result.String())
	return invitationInfo, err
}
