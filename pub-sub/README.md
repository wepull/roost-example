# Pub-Sub Example

Pub-Sub application demostrates zbio messaging platform.
Producer creates a topic in zbio and messages are sent to the topic.
Consumer consumes messages from that topic.

## Commands to build and deploy Pub Sub app in k8s cluster

```bash
# Build application
make build

## Dockerise
make docker

## If using kind cluster
kind load docker-image zbio-example/consumer:v1
kind load docker-image zbio-example/producer:v1

# Deploy
make deploy

# View producer and consumer's log
kubectl logs zbio-sample-producer
kubectl logs zbio-sample-consumer
```

### Cleaning

```bash
make clean
```

- _Deletes producer and consumer binaries_
- _Deletes deployed pub sub app resources from kubernetes_
