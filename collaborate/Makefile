
all: build dockerise

build:
	CGO_ENABLED=0 GOOS=linux go build -o fetcher cmd/fetcher/main.go
	CGO_ENABLED=0 GOOS=linux go build -o retriever cmd/retriever/main.go

dockerise: build
	@docker build -f fetcher.Dockerfile -t roost-example/fetcher:v1 .
	@docker build -f retriever.Dockerfile -t roost-example/retriever:v1 .
	${MAKE} clean

deploy:
	@kubectl apply -f kubernetes-manifests/fetcher.yaml 
	@kubectl apply -f kubernetes-manifests/retriever.yaml 

undeploy:
	@kubectl delete -f kubernetes-manifests/fetcher.yaml >/dev/null 2>&1 || true
	@kubectl delete -f kubernetes-manifests/retriever.yaml >/dev/null 2>&1 || true

clean:
	rm -f fetcher retriever

