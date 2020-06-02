
all: clean build dockerise

build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/fetcher cmd/fetcher/main.go
	CGO_ENABLED=0 GOOS=linux go build -o bin/retriever cmd/retriever/main.go

dockerise: build
	@docker build -f dockerfiles/fetcher.Dockerfile -t roost-example/fetcher:v1 bin
	@docker build -f dockerfiles/retriever.Dockerfile -t roost-example/retriever:v1 bin
	${MAKE} clean

deploy:
	@kubectl apply -f kubernetes-manifests/fetcher.yaml 
	@kubectl apply -f kubernetes-manifests/retriever.yaml 

undeploy:
	@kubectl delete -f kubernetes-manifests/fetcher.yaml >/dev/null 2>&1 || true
	@kubectl delete -f kubernetes-manifests/retriever.yaml >/dev/null 2>&1 || true

clean:
	-@rm -rf bin/*
	-${MAKE} undeploy

