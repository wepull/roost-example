#MAKE fike for implementing example

.PHONY: all
all: clean build dockerise deploy
	$(MAKE) clean_bin

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

.PHONY: dockerise
dockerise:
	DOCKER_BUILDKIT=1 docker build -f make.Dockerfile -t html:make_scratch .

# Dockerise app with golang as base image
build-image-html-golang:
	DOCKER_BUILDKIT=1 docker build -t html:golang .

# Dockerise app with multistage build and alpine as base image
build-image-html-alpine:
	DOCKER_BUILDKIT=1 docker build -f alpine.Dockerfile -t html:alpine .

# Dockerise app with multistage build and scratch as base image
build-image-html-scratch:
	DOCKER_BUILDKIT=1 docker build -f scratch.Dockerfile -t html:scratch .

# Run the html:golang image
docker-run-html-golang:
	docker run -p 30047:8080 html:golang

# Run the html:alpine image
docker-run-html-alpine:
	docker run -p 30047:8080 html:alpine

# Run the html:scratch image
docker-run-html-scratch:
	docker run -p 30047:8080 html:scratch

# Deploy in the kubernetes cluster
deploy:
	kubectl apply -f k8s-deployment.yaml

# Clean 
clean: undeploy

.PHONY: undeploy
undeploy:
	kubectl delete -f k8s-deployment.yaml >/dev/null 2>&1 || true

.PHONY: clean_bin
clean_bin:
	rm -f app