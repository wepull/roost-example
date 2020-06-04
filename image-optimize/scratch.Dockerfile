FROM golang:alpine as builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT [ "/app/main" ]