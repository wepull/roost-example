APP=producer
APP_VERSION=v1

build:
	export GOPRIVATE="github.com/ZB-io/*"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o bin/${APP}

.PHONY: docker
docker:
	docker build -t zbio-example/${APP}:${APP_VERSION} .

clean:
	rm -f bin/${APP}
