FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD main.go  /go/src/github.com/roost-io/roost-example/collaborate/retriever/
WORKDIR /go/src/github.com/roost-io/roost-example/collaborate/retriever
RUN CGO_ENABLED=0 GOOS=linux go build -o retriever main.go

FROM alpine:3.9
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY --from=builder /go/src/github.com/roost-io/roost-example/collaborate/retriever/retriever /app/
EXPOSE 8081
ENTRYPOINT [ "/app/retriever" ]