FROM golang:latest
WORKDIR /app
COPY main.go .
COPY *.html .
RUN go build -o main
EXPOSE 8080
CMD [ "./main" ]

