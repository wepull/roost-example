APP ?= googlebookapi
IMAGE_VERSION ?= v1
REGISTRY ?= localhost:5000
#IMAGE ?= ${REGISTRY}/zbio-example/${APP}:${IMAGE_VERSION}
IMAGE ?= zbio-example/${APP}:${IMAGE_VERSION}

.PHONY: all
all: clean build dockerise deploy

.PHONY: build
build:	
	export GOPRIVATE="github.com/ZB-io/*"
	cd src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o ${APP}

.PHONY: deploy
deploy:
	#kubectl delete -f src/googlebookapi.yaml --now
	kubectl apply -f src/googlebookapi.yaml

.PHONY: clean
clean:
	rm -f src/googlebookapi

.PHONY: dockerise
dockerise:
	/usr/bin/docker ps >/dev/null 2>&1 && docker build -f src/Dockerfile -t ${IMAGE} src
	../push_image.sh "${IMAGE}"

.PHONY: release
release: 
	$(MAKE) build
	$(MAKE) dockerise
	$(MAKE) deploy
	
	
		
