all: 
	export GOPRIVATE="github.com/ZB-io/*"
	GOOS=linux GOARCH=amd64 go build -o bin/zbsample_consumer
	cp ../../security/cert/server.crt bin/
ifdef NODOCKER
	echo "DO NOT BUILD IMAGE"
else 
	docker build -f docker/Dockerfile -t zbio-example/consumer:v1 .
endif

clean:
	rm -f bin/zbsample_consumer
