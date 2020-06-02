# Dependent microservices for collaborate feature in ROOST

Two containerised services `fetcher` and `retriver` run in kubernetes environment in two pods. `fetcher` services expose an api at <http://roost-master:30047/articles?tag=kubernetes> to fetch latest articles from <http://dev.to> based on provided tags and stores in provisioned pod volume. `retriver` service is dependent on `fetcher` service to retrieve stored articles in `fetcher` service volume. As soon as articles file is found, `retriever` service reads the content and sends output over HTTP/browser.

## Commands to deploy

```bash
# Single command to build, dockerise and deploy application in kubernetes.
make

# Builds applicaiton binaries
make build

# Dockerise: Build docker images
make dockerise
```

## How to test

1. Check if both containers are running in kubernetes

    ```bash
        kubectl get pod collaborate
    ```

2. `fetcher` service must be running

    ```bash
        kubectl logs collab-fetcher
    ```

3. Above content must be served by `retriever` service.

    ```bash
        curl "http://roost-master:30048/articles?tag=kubernetes"
        # --tail n: shows last n lines from logs
        kubectl logs collab-retriever --tail 400
    ```
