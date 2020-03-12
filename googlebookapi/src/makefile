APP=googlebookapi
.PHONY: all
all: 
	export GOPRIVATE="github.com/ZB-io/*"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${APP}
ifdef NODOCKER
	echo "DO NOT BUILD IMAGE"
else 
	docker build -f docker/Dockerfile -t zbio-example/${APP}/src:v1 .
endif

.PHONY: deploy
deploy:
	kubectl apply -f googlebookapi.yaml

.PHONY: clean
clean:
	rm -f bin/googlebookapi
	kubectl delete -f googlebookapi.yaml --all --now=true 
