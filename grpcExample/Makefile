# Copyright 2019 American Express Travel Related Services Company, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
# in compliance with the License. You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under the License
# is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
# or implied. See the License for the specific language governing permissions and limitations under
# the License.

CLIENT_IMG=clientimg
SERVER_IMG=serverimg
IMAGE_TAG=v1

GRPC_HEALTH_PROBE=grpc_health_probe-linux-amd64
GRPC_HEALTH_PROBE_PATH=../tool/${GRPC_HEALTH_PROBE}

HOSTNAME := $(shell hostname)

.PHONY: all
all: clean build dockerise deploy
	$(MAKE) clean_bin

.PHONY: build
build: build-server build-client

.PHONY: build-server
build-server:
	cd server-grpc && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -gcflags='-N -l' -o ../bin/grpc-server && cd -

.PHONY: build-client
build-client:
	cd client-grpc && GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' -o ../bin/grpc-client && cd -

.PHONY: dockerise
dockerise:
	docker build -t ${CLIENT_IMG}:${IMAGE_TAG} -f client.Dockerfile .
	docker build -t ${SERVER_IMG}:${IMAGE_TAG} -f server.Dockerfile .	
	${shell pwd}/../push_image.sh "${CLIENT_IMG}:${IMAGE_TAG}"
	${shell pwd}/../push_image.sh "${SERVER_IMG}:${IMAGE_TAG}"

.PHONY: deploy
deploy:
	kubectl apply -f kubernetes/deploy.yaml

.PHONY: clean
clean: clean_bin
	kubectl delete -f kubernetes/deploy.yaml --now >/dev/null 2>&1 || true

.PHONY: clean_bin
clean_bin:
	rm -f bin/*
