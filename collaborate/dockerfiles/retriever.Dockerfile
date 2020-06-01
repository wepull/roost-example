FROM alpine
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes"
COPY retriever /app/retriever
ENTRYPOINT [ "/app/retriever" ]