# Pub-Sub Example

Pub-Sub application demostrates zbio messaging platform.
Producer creates a topic in zbio and messages are sent to the topic.
Consumer consumes messages from that topic.

## Commands to deploy in ZKE Cluster from Roost

```bash
# Build, Dockerise and deploy into ZKE cluster
make

# View producer and consumer's log
kubectl logs zbio-sample-producer
kubectl logs zbio-sample-consumer

# View zbio service logs
kubectl logs service/zbio-service --namespace zbio
```

### Cleaning

```bash
make clean
```

- _Deletes producer and consumer binaries_
- _Deletes deployed pub sub app resources from kubernetes_
