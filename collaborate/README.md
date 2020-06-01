# Dependent microservices for collaborate feature in ROOST

Two containerised services `fetcher` and `retriver` run in kubernetes environment in single pod but in two containers. `fetcher` services expose an api at <http://roost-master:30047/articles?tag=kubernetes> to fetch latest articles from <http://dev.to> and stores in provisioned pod volume. `retriver` service is dependent on `fetcher` service to create file where fetched articles are stored. As soon as articles file is found, `retriever` service reads the content and sends output over HTTP/browser.

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
        kubectl logs pod collaborate fetcher
    ```

3. Make `curl` request to fetcher service to retrive articles which save articles into file

    ```bash
        curl "http://roost-master:30047/articles?tag=kubernetes"
        kubectl logs pod collaborate fetcher
    ```

4. Above content must be served by `retriever` service.

    ```bash
        curl "http://roost-master:30048/"
        # --tail n: shows last n lines from logs
        kubectl logs collaborate retriever --tail 400
    ```
