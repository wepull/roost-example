FROM scratch
COPY app view.html /app/
EXPOSE 8080
ENTRYPOINT [ "/app/app" ]