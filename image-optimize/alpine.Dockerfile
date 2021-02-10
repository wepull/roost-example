FROM golang:alpine
WORKDIR /app
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
ENTRYPOINT [ "/app/main" ]
