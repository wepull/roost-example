#MAKE fike for implementing example

.PHONY: all
all: clean golang alpine scratch roost
	
run: run-golang run-alpine run-scratch run-roost

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Dockerise app with golang as base image
golang:
	docker build -f golang.Dockerfile -t image-optimize:golang .

# Dockerise app with multistage build and alpine as base image
alpine:
	docker build -f alpine.Dockerfile -t image-optimize:alpine .

# Dockerise app with multistage build and scratch as base image
scratch:
	docker build -f scratch.Dockerfile -t image-optimize:scratch .
	
# Dockerise app with multistage build and scratch as base image
roost: build
	docker build -f roost.Dockerfile -t image-optimize:roost .
	
# Dockerise app with golang as base image
run-golang:
	docker run image-optimize:golang /app/main

# Dockerise app with multistage build and alpine as base image
run-alpine:
	docker run image-optimize:alpine /app/main

# Dockerise app with multistage build and scratch as base image
run-scratch:
	docker run image-optimize:scratch

run-roost:
	docker run image-optimize:roost

clean:
	docker images | grep 'image-optimize' | awk '{print $$3}' | xargs -r docker rmi -f
	rm -f app
