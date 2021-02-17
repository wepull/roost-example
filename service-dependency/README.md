# Service Fitness example

## How service fitness works ?

### Deploy application

```bash
  kubectl apply -f kubernetes-manifests/
```

### Deploy application service test suite (this is dummy test suite)

1. build the container image [service-test-suite](./service-test-suite)
2. deploy test suite into ZKE [test-suite.yaml](./service-test-suite/test-suite.yaml)

### Upload service-dependency json

upload [service-dependency.json](./service-dependency.json) into Observability -> Service Fitness and apply to cluster.

### Observe impacted services in event viewer during collaboration 

  1. Impact due to busybox image (as currencyservice depends on `busybox` image)
  
```bash
    docker pull busybox:latest
```

  2. impact due to deployment .yaml file received. (any yaml file name mentioned in service-dependency.json)

