FROM golang:1.12-alpine as builder
ENV PROJECT github.com/roost-io/roost-example/grpc-k8s-health-check
WORKDIR /go/src/$PROJECT
COPY . .
WORKDIR /go/src/$PROJECT/server-grpc/
RUN go build -gcflags='-N -l' -o /app

FROM alpine:3.9
# Get dependencies
RUN apk add wget
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 
# Make it executable 
RUN chmod +x /bin/grpc_health_probe  

COPY --from=builder /app /app
# Expose port to the outside once the container has launched
EXPOSE 3000
# The ENTRYPOINT of an image specifies what executable to run when the container starts
ENTRYPOINT [ "/app" ]
