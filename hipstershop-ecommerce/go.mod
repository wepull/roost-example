module github.com/roost-io/roost-example/hipstershop-ecommerce

go 1.14

require (
	cloud.google.com/go v0.40.0
	contrib.go.opencensus.io/exporter/stackdriver v0.9.1
	git.apache.org/thrift.git v0.0.0-20180807212849-6e67faa92827 // indirect
	github.com/Azure/azure-sdk-for-go v30.1.0+incompatible // indirect
	github.com/Azure/go-autorest v11.1.2+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.9.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.2.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.1.0 // indirect
	github.com/GoogleCloudPlatform/microservices-demo v0.2.0
	github.com/Shopify/sarama v1.24.1 // indirect
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/ZB-io/zbio v0.0.6
	github.com/apache/thrift v0.13.0 // indirect
	github.com/aws/aws-sdk-go v1.23.20 // indirect
	github.com/bsm/sarama-cluster v2.1.15+incompatible // indirect
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f // indirect
	github.com/dghubble/oauth1 v0.6.0 // indirect
	github.com/dgraph-io/ristretto v0.0.1 // indirect
	github.com/eko/gocache v1.0.0 // indirect
	github.com/fionita/linkedin-go v0.0.0-20170522061944-de3395609bba // indirect
	github.com/go-gl/glfw v0.0.0-20190409004039-e6da0acd62b1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/mock v1.4.3 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.4.1
	github.com/google/martian v2.1.0+incompatible // indirect
	github.com/google/pprof v0.0.0-20200430221834-fc25d7d30c6d // indirect
	github.com/google/uuid v1.1.1
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/gophercloud/gophercloud v0.1.0 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.8.5 // indirect
	github.com/hashicorp/go-memdb v1.0.4 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/influxdata/influxdb1-client v0.0.0-20191209144304-8bf82d3c094d // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/cpuid v1.2.2 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lightstep/tracecontext.go v0.0.0-20181129014701-1757c391b1ac // indirect
	github.com/mattn/go-sqlite3 v2.0.2+incompatible // indirect
	github.com/nats-io/nats-server/v2 v2.1.2 // indirect
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20191001143057-db30781987df // indirect
	github.com/openzipkin/zipkin-go v0.1.6 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/quipo/statsd v0.0.0-20180118161217-3d6a5565f314 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	go.opencensus.io v0.22.0
	go.opentelemetry.io/otel v0.3.0 // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200204104054-c9f3fb736b72 // indirect
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20200501052902-10377860bb8e // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20200501065659-ab2804fb9c9d // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200430143042-b979b6f78d84 // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/confluentinc/confluent-kafka-go.v1 v1.1.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
	k8s.io/api v0.17.3 // indirect
	k8s.io/utils v0.0.0-20200108110541-e2fb8e668047 // indirect
	pack.ag/amqp v0.11.0 // indirect
	rsc.io/binaryregexp v0.2.0 // indirect
)

replace (
	go.opencensus.io => go.opencensus.io v0.16.0
	contrib.go.opencensus.io/exporter/stackdriver => contrib.go.opencensus.io/exporter/stackdriver v0.5.0
)