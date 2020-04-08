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

#FROM ubuntu:18.04
#FROM golang:latest
#RUN apt update
#RUN apt install -y git

#RUN git --version

FROM golang:1.12.5-alpine3.9 AS BUILD_STAGE
#RUN apt-get update
#RUN apt-get install gcc make
#RUN add-apt-repository ppa:duh/golang
#RUN apt-get install golang

#RUN apt-get update && yes | apt-get upgrade
#FROM ubuntu
# Add vendor to GOPATH
ADD vendor /go/src
# Add client-grpc to the docker image
ADD client-grpc /go/src/client-server-grpc/client-grpc
COPY Collect_Data.go /go/src/client-server-grpc/client-grpc

COPY cert /go/src/client-server-grpc/client-grpc
ENV SERVER_CRT=/go/src/client-server-grpc/client-grpc

RUN ls /go/src/github.com/ZB-io/zbio/security/cert/
#RUN go get -u github.com/ZB-io/zbio/client
#RUN export GO111MODULE=on
#RUN go get github.com/ZB-io/zbio/client
# Add api to to the docker image
ADD api /go/src/client-server-grpc/api
# Set working dir in the container 
WORKDIR /go/src/client-server-grpc/client-grpc
#RUN apt-get update && yes | apt-get upgrade
#RUN apt-get install git 

RUN ls /go/src/client-server-grpc/client-grpc
# Build the program and output it to app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app
RUN ls -al /app
# Make the dockerfile more optimized by using multistage dockerbuild which we copy the binary from the BUILD_STAGE container to the final container.
# The second FROM instruction starts a new build stage with the alpine image as its base.
FROM alpine:3.9
RUN ls
COPY --from=BUILD_STAGE /app /app
# The ENTRYPOINT of an image specifies what executable to run when the container starts.
ENTRYPOINT ["/app"]
