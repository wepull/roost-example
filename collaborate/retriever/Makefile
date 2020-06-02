all: build dockerise

build:
	CGO_ENABLED=0 GOOS=linux go build -o retriever main.go

dockerise: build
	@docker build -f Dockerfile -t retriever:latest .
	${shell pwd}/../../push_image.sh retriever:latest

deploy:
	@kubectl apply -f retriever.yaml 

undeploy:
	@kubectl delete -f retriever.yaml >/dev/null 2>&1 || true

clean:
	rm -f retriever
