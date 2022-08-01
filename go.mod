module Open_IM

go 1.15

require (
	cloud.google.com/go/firestore v1.6.1 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/Shopify/sarama v1.32.0
	github.com/alibabacloud-go/darabonba-openapi v0.1.11
	github.com/alibabacloud-go/dysmsapi-20170525/v2 v2.0.8
	github.com/alibabacloud-go/sts-20150401 v1.1.0
	github.com/alibabacloud-go/tea v1.1.17
	github.com/antonfisher/nested-logrus-formatter v1.3.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dtm-labs/rockscache v0.0.11
	github.com/fatih/structs v1.1.0
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/validator/v10 v10.11.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gogo/protobuf v1.3.2
	github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/copier v0.3.4
	github.com/jinzhu/gorm v1.9.16
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/minio/minio-go/v7 v7.0.22
	github.com/mitchellh/mapstructure v1.4.2
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/olivere/elastic/v7 v7.0.23
	github.com/pkg/errors v0.9.1
	github.com/razeencheng/demo-go/swaggo-gin v0.0.0-20220621231537-0a4e9bb56373
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.2
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe
	github.com/swaggo/gin-swagger v1.5.0
	github.com/swaggo/swag v1.8.4
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.428
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms v1.0.428
	github.com/tencentyun/qcloud-cos-sts-sdk v0.0.0-20210325043845-84a0811633ca
	go.etcd.io/etcd/api/v3 v3.5.4
	go.etcd.io/etcd/client/v3 v3.5.4
	go.mongodb.org/mongo-driver v1.8.3
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/api v0.59.0
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

replace github.com/Shopify/sarama => github.com/Shopify/sarama v1.29.0
