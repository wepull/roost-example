FROM scratch
COPY checkoutservice /checkoutservice
COPY grpc_health_probe-linux-amd64 /bin/grpc_health_probe
EXPOSE 5050
ENTRYPOINT ["/checkoutservice"]
