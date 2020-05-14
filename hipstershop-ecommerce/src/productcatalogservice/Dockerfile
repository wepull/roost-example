#FROM golang:1.12-alpine AS builder
#RUN apk add --no-cache ca-certificates git && \
#      wget -qO/go/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
#      chmod +x /go/bin/dep

#ENV PROJECT github.com/GoogleCloudPlatform/microservices-demo/src/productcatalogservice
#WORKDIR /go/src/$PROJECT

# restore dependencies
#COPY Gopkg.* ./
#RUN dep ensure --vendor-only -v

#COPY . .
#RUN go build -o /productcatalogservice .

FROM alpine AS release
RUN apk add --no-cache ca-certificates
COPY grpc_health_probe-linux-amd64 /bin/grpc_health_probe
WORKDIR /productcatalogservice
#COPY --from=builder /productcatalogservice ./server
COPY productcatalogservice ./server
COPY products.json .
EXPOSE 3550
ENTRYPOINT ["/productcatalogservice/server"]

