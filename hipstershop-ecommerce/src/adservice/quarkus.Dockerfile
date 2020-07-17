FROM openjdk:8-slim as builder

WORKDIR /app

COPY . .
RUN chmod +x gradlew