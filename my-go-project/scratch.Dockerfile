FROM golang:latest as builder
WORKDIR /app
COPY main.go .
COPY *.html .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT [ "/app/main" ]