FROM scratch
COPY grpc_health_probe-linux-amd64 /bin/grpc_health_probe
WORKDIR /productcatalogservice
COPY productcatalogservice ./server
COPY products.json .
EXPOSE 3550
ENTRYPOINT ["/productcatalogservice/server"]

