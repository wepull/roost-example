FROM golang:alpine
MAINTAINER VIVEK K B <kbvivek@zb.io>
LABEL Description="This image is used to start producer " Vendor="Zettabytes" Version="1.0"

# ADD producer binary location
ADD bin /producer
WORKDIR /producer
ENV SERVER_CRT=/producer/server.crt
ENTRYPOINT ["./producer"]
