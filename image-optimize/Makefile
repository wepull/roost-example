#MAKE fike for implementing example

.PHONY: all
all: clean build dockerise deploy
	$(MAKE) clean_bin

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Dockerise app with golang as base image
golang:
	DOCKER_BUILDKIT=1 docker build -f golang.Dockerfile -t image-optimize:golang .

# Dockerise app with multistage build and alpine as base image
alpine:
	DOCKER_BUILDKIT=1 docker build -f alpine.Dockerfile -t image-optimize:alpine .

# Dockerise app with multistage build and scratch as base image
scratch:
	DOCKER_BUILDKIT=1 docker build -f scratch.Dockerfile -t image-optimize:scratch .
	
# Dockerise app with multistage build and scratch as base image
roost:
	DOCKER_BUILDKIT=1 docker build -f roost.Dockerfile -t image-optimize:roost .
	
# Dockerise app with golang as base image
run-golang:
	DOCKER_BUILDKIT=1 docker run image-optimize:golang /app/main

# Dockerise app with multistage build and alpine as base image
run-alpine:
	DOCKER_BUILDKIT=1 docker run image-optimize:alpine /app/main

# Dockerise app with multistage build and scratch as base image
run-scratch:
	DOCKER_BUILDKIT=1 docker run image-optimize:scratch

run-roost:
	DOCKER_BUILDKIT=1 docker run image-optimize:roost

.PHONY: clean_bin
clean_bin:
	rm -f app