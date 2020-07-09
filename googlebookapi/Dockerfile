FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD src /go/src/github.com/roost-io/roost-example/googlebookapi/src
WORKDIR /go/src/github.com/roost-io/roost-example/googlebookapi/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -o googlebookapi

FROM scratch
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY  --from=builder /go/src/github.com/roost-io/roost-example/googlebookapi/src/googlebookapi /app/
EXPOSE 9999
ENTRYPOINT ["/app/googlebookapi"]
