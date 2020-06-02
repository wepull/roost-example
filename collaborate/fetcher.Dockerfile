FROM alpine
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes"
COPY fetcher /app/fetcher
EXPOSE 8080
ENTRYPOINT [ "/app/fetcher" ]