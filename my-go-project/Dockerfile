FROM golang:latest
WORKDIR /app
COPY main.go .
COPY *.html .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
EXPOSE 8080
CMD [ "./main" ]

