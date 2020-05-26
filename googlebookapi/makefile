APP ?= googlebookapi
IMAGE_VERSION ?= v1
IMAGE ?= zbio-example/${APP}:${IMAGE_VERSION}
HOSTNAME=$(shell hostname)

.PHONY: all
all: clean build dockerise deploy
	$(MAKE) clean_bin

.PHONY: build
build:	
	export GOPRIVATE="github.com/ZB-io/*"
	cd src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o ${APP}

.PHONY: deploy
deploy:
	kubectl apply -f src/googlebookapi.yaml

.PHONY: dockerise
dockerise:
	docker build -f src/Dockerfile -t ${IMAGE} src
	${shell pwd}/../push_image.sh "${IMAGE}"

.PHONY: clean
clean: clean_bin
	kubectl delete -f src/googlebookapi.yaml --now >/dev/null 2>&1 || true

.PHONY: clean_bin
clean_bin:
	rm -f googlebookapi
		
