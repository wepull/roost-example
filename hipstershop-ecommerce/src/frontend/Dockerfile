FROM scratch
WORKDIR /frontend
COPY frontend /frontend/server
COPY ./templates ./templates
COPY ./static ./static
EXPOSE 8080
ENTRYPOINT ["/frontend/server"]
