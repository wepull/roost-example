FROM mgdevstack/grpc-go:protobuf-v3.11.2 as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD . /go/src/github.com/roost-io/roost-example/grpcExample/
WORKDIR /go/src/github.com/roost-io/roost-example/grpcExample/server-grpc
RUN GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' -o ../bin/grpc-server

FROM alpine:3.9
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY --from=builder /go/src/github.com/roost-io/roost-example/grpcExample/grpc_health_probe-linux-amd64 /bin/grpc_health_probe
COPY --from=builder /go/src/github.com/roost-io/roost-example/grpcExample/bin/grpc-server /app/grpc-server
USER nobody
ENTRYPOINT [ "/app/grpc-server" ]