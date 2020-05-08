FROM alpine:3.9
COPY bin/grpc-server /app/grpc-server
COPY grpc_health_probe-linux-amd64 /bin/grpc_health_probe
EXPOSE 3000
ENTRYPOINT [ "/app/grpc-server" ]
