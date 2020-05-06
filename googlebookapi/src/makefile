APP ?= googlebookapi
REGISTRY ?= localhost:5000
IMAGE_VERSION ?= v1
IMAGE ?= ${REGISTRY}/zbio-example/${APP}:${IMAGE_VERSION}

.PHONY: all
all: 
	export GOPRIVATE="github.com/ZB-io/*"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o bin/${APP}
ifdef NODOCKER
	echo "DO NOT BUILD IMAGE"
else 
	docker build -f Dockerfile -t zbio-example/${APP}:v1 bin
endif

.PHONY: build
build:	
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o bin/${APP}

.PHONY: deploy
deploy:
	kubectl apply -f googlebookapi.yaml

.PHONY: clean
clean:
	rm -f bin/googlebookapi
	kubectl delete -f googlebookapi.yaml --now

.PHONY: dockerise
dockerise:
	docker build -f Dockerfile -t ${IMAGE} bin
	docker push ${IMAGE}

.PHONY: release
release: 
	$(MAKE) build
	$(MAKE) dockerise
	$(MAKE) deploy
	
	
		
