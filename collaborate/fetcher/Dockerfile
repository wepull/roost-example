FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD main.go  /go/src/github.com/roost-io/roost-example/collaborate/fetcher/
WORKDIR /go/src/github.com/roost-io/roost-example/collaborate/fetcher
RUN CGO_ENABLED=0 GOOS=linux go build -o fetcher main.go

FROM alpine:3.9
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY --from=builder /go/src/github.com/roost-io/roost-example/collaborate/fetcher/fetcher /app/
EXPOSE 8080
ENTRYPOINT [ "/app/fetcher" ]