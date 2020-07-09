all:  dockerise undeploy deploy
	${MAKE} clean

build:
	CGO_ENABLED=0 GOOS=linux go build -o retriever main.go

dockerise:
	@docker build -f Dockerfile -t retriever:latest .
	${shell pwd}/../../push_image.sh retriever:latest

deploy:
	@kubectl apply -f retriever.yaml 

.PHONY: undeploy
undeploy:
	@kubectl delete -f retriever.yaml >/dev/null 2>&1 || true

clean:
	- rm -f retriever
