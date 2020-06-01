FROM alpine
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes"
COPY fetcher /app/fetcher
ENTRYPOINT [ "/app/fetcher" ]