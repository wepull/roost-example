# Googlebookapi Example

## Commands to build and deploy Googlebookapi app in k8s cluster

### How to use in localhost
* It runs on localhost:9999/books/<isbn_number>
* ISBN stands for international standard book number , which is
    13 digit number uniquely identify all the books.
* ISBN can be found in internet (eg: 978-1788999786 (Name: Mastering Kubernetes))
* The digit should not have any special characters in between.

### When using Kind cluster
- _build images need to be loaded into the kind cluster_

```bash
make
kind load docker-image zbio-sample-googlebookapi
make deploy
```

- _Rest Trigger will be listening on http://0.0.0.0:9999_ 

### Cleaning

```bash
make clean
```

- _Deletes googlebookapi binaries_
- _Deletes deployed googlebookapi app resources from kubernetes_
