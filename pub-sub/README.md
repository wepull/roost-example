# Pub Sub Example

## Commands to build and deploy Pub Sub app in k8s cluster

### When using Docker Desktop

```bash
make
```

### When using Kind cluster
- _build images need to be loaded into the kind cluster_

```bash
make build
make docker
kind load docker-image zbio-example/consumer:v1
kind load docker-image zbio-example/producer:v1
make deploy
```

### Cleaning

```bash
make clean
```

- _Deletes producer and consumer binaries_
- _Deletes deployed pub sub app resources from kubernetes_
