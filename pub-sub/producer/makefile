APP=producer
APP_VERSION=v1
IMAGE="zbio-example/${APP}:${APP_VERSION}"
HOSTNAME := $(shell hostname)

build:
	export GOPRIVATE="github.com/ZB-io/*"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o bin/${APP}

.PHONY: docker
docker:
	docker ps >/dev/null 2>&1 && docker build -t ${IMAGE} .

clean:
	rm -f bin/${APP}
