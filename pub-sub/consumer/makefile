APP=consumer
APP_VERSION=v1
IMAGE="zbio-example/${APP}:${APP_VERSION}"
HOSTNAME := $(shell hostname)

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o bin/${APP}

.PHONY: docker
docker:
	docker build -t ${IMAGE} .
ifeq ($(HOSTNAME), roost-utility)
		${shell pwd}/../../push_image.sh "${IMAGE}"
endif

clean:
	rm -f bin/${APP}
