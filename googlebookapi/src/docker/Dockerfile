FROM alpine:latest
ADD bin /app
ADD cert /app
ENV SERVER_CRT="/app/server.crt"
EXPOSE 9999
ENTRYPOINT ["/app/googlebookapi"]
