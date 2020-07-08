FROM golang:alpine as builder
LABEL maintainer="Zettabytes"
ADD src /go/src/github.com/roost-io/roost-example/googlebookapi/src
WORKDIR /go/src/github.com/roost-io/roost-example/googlebookapi/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o googlebookapi

FROM scratch
COPY  --from=builder /go/src/github.com/roost-io/roost-example/googlebookapi/src/googlebookapi /app/
EXPOSE 9999
ENTRYPOINT ["/app/googlebookapi"]
