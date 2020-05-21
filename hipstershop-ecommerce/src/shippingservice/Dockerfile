FROM scratch
COPY grpc_health_probe-linux-amd64 /bin/grpc_health_probe
COPY shippingservice /shippingservice
ENV APP_PORT=50051
EXPOSE 50051
ENTRYPOINT ["/shippingservice"]
