FROM zbbase:alpine
MAINTAINER VIVEK K B <kbvivek@zb.io>
LABEL Description="This image is used to start producer " Vendor="Zettabytes" Version="1.0"

# ADD producer binary location
ADD bin /zbsample
WORKDIR /zbsample

ENV SERVER_CRT="/zbsample/server.crt"
#ENTRYPOINT ["/zbsample/producer"]

ENTRYPOINT ./producer.sh
CMD [default] 
