APP=producer
APP_VERSION=v1
IMAGE=${APP}:${APP_VERSION}

build:
	export GOPRIVATE="github.com/ZB-io/*"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o bin/${APP}

.PHONY: docker
docker:
ifeq ($(HOSTNAME), roost-utility)
	docker ps >/dev/null 2>&1 && docker build -t zbio-example/${IMAGE} .
	@sh ${shell pwd}/../../push_image.sh "${IMAGE}"
else
	docker build -t zbio-example/${IMAGE} .
endif

clean:
	rm -f bin/${APP}
