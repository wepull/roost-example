APP ?= googlebookapi
IMAGE_VERSION ?= v1
IMAGE ?= zbio-example/${APP}:${IMAGE_VERSION}
HOSTNAME=$(shell hostname)

.PHONY: all
all: clean build dockerise deploy

.PHONY: build
build:	
	export GOPRIVATE="github.com/ZB-io/*"
	cd src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o ${APP}

.PHONY: deploy
deploy:
	kubectl apply -f src/googlebookapi.yaml

.PHONY: clean
clean:
	rm -f src/googlebookapi
	kubectl delete -f src/googlebookapi.yaml --now || true

.PHONY: dockerise
dockerise:
ifeq ($(HOSTNAME), roost-utility)
	/usr/bin/docker ps >/dev/null 2>&1 && docker build -f src/Dockerfile -t ${IMAGE} src
	${shell pwd}/../push_image.sh "${IMAGE}"
else
	@docker build -f src/Dockerfile -t ${IMAGE} src
endif
		
