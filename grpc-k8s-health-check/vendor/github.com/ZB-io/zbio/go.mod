module github.com/ZB-io/zbio

go 1.13

require (
	github.com/Shopify/sarama v1.24.1
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/allegro/bigcache v1.2.1
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/cloudevents/sdk-go v1.1.1
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f
	github.com/dghubble/oauth1 v0.6.0
	github.com/dgraph-io/ristretto v0.0.1
	github.com/eko/gocache v1.0.0
	github.com/fionita/linkedin-go v0.0.0-20170522061944-de3395609bba
	github.com/go-kit/kit v0.9.0
	github.com/golang/protobuf v1.3.3
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/go-memdb v1.0.4
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/influxdata/influxdb1-client v0.0.0-20191209144304-8bf82d3c094d
	github.com/klauspost/cpuid v1.2.2 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
	github.com/mitchellh/go-homedir v1.1.0
	github.com/onsi/gomega v1.8.1
	github.com/opentracing-contrib/go-grpc v0.0.0-20191001143057-db30781987df
	github.com/opentracing/opentracing-go v1.1.1-0.20190913142402-a7454ce5950e
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/common v0.9.1
	github.com/quipo/statsd v0.0.0-20180118161217-3d6a5565f314
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.4.0
	go.opentelemetry.io/otel v0.3.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.3.0
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/genproto v0.0.0-20200212174721-66ed5ce911ce // indirect
	google.golang.org/grpc v1.27.1
	gopkg.in/confluentinc/confluent-kafka-go.v1 v1.1.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.17.3
	k8s.io/apimachinery v0.17.3
	k8s.io/client-go v0.17.3
	k8s.io/utils v0.0.0-20200108110541-e2fb8e668047 // indirect
)

replace (
	github.com/ZB-io/zbio => ./
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
)
