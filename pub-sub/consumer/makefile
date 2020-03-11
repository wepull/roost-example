APP=consumer
APP_VERSION=v1

build:
	export GOPRIVATE="github.com/ZB-io/*"
	GOOS=linux GOARCH=amd64 go build -o bin/${APP}

.PHONY: docker
docker:
	docker build -t zbio-example/${APP}:${APP_VERSION} .

clean:
	rm -f bin/${APP}
