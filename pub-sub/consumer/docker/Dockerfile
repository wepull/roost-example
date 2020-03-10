FROM golang:alpine
MAINTAINER VIVEK K B <kbvivek@zb.io>
LABEL Description="This image is used to start consumer " Vendor="Zettabytes" Version="1.0"

# ADD consumer binary location
ADD bin /consumer
WORKDIR /consumer

ENV SERVER_CRT="/consumer/server.crt"

ENTRYPOINT ./consumer.sh
CMD [default]

