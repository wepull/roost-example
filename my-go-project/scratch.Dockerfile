FROM golang:latest as builder
WORKDIR /app
COPY main.go .
COPY *.html .
RUN go build -o main

FROM scratch
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT [ "/app/main" ]
