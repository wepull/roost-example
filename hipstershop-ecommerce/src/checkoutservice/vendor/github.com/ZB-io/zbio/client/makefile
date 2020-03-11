
REPO=..
include ${REPO}/common.mk

.PHONY: all
all: go-fmt
	make build-linux

.PHONY: build
build:
	CGO_ENABLED=0 go generate
	go build
	CGO_ENABLED=0 go test -c

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test -c

.PHONY: clean
clean:
	rm -rf docker/zbclient client.test
	find ../rpc/common -name '*.pb.go' -exec rm {} \;
	go clean

.PHONY: image
image:
	echo "Building ZBClient"
	mkdir -p docker/zbclient/security/cert
	cp client.test zbclient.sh docker/zbclient
	cp ../security/cert/server.crt docker/zbclient/security/cert
	#docker build -f Dockerfile -t zbio/zbclient:v1 .
	make -C docker
