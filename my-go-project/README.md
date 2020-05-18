# HTTP based application to demostrate use of optimised Dockerfiles

## How to run this application

###### Using Roost Desktop Engine (RDE)

> Right-click on `Makefile` and click `Run` for hassle-free deployment to ZKE

> What all is done by `make`?

* Cleans existing deployment [k8s-deployment.yaml](./k8s-deployment.yaml)

* Removes the application binary (if exists) [app](app)

* Builds an executable by compiling this codebase [main.go](./main.go)

* Builds an image using the app binary [make.Dockerfile](./make.Dockerfile)

* Deploys the image to ZKE Cluster [k8s-deployment.yaml](./k8s-deployment.yaml)

## View application logs

> Using `Workload Analytics` (RDE) for deployed application

> [RDE Workload Analytics image](show_GoogleBookAPI_pod_logs_and_workload_view) (to be added)

> Using RKT Konsole

```bash
kubectl logs service/myapp
```

## How to access app

> Using any browser

* Open <http://roost-master:30047>

### Following are valid `make` commands

Change working directory to `$GOPATH/src/github.com/roost-io/roost-example/` to run following valid `make` commands.

```bash
cd $GOPATH/src/github.com/roost-io/roost-example/
```

###### Commands are -

1. Dockerise app with `golang` as base image

```bash
make build-image-html-golang
```

1. Dockerise app with `multistage build and alpine` as base image

```bash
make build-image-html-alpine
```

1. Dockerise app with `multistage build and scratch` as base image

```bash
make build-image-html-scratch
```

1. Run application with `html:golang` image (accessible at [http://roost-master:30047](http://roost-master:30047))

```bash
make docker-run-html-golang
```

1. Run application with `html:alpine` image (accessible at [http://roost-master:30047](http://roost-master:30047))

```bash
make docker-run-html-alpine
```

1. Run application with `html:scratch` image (accessible at [http://roost-master:30047](http://roost-master:30047))

```bash
make docker-run-html-scratch
```

1. Clean [undeploy app from kubernetes]

```bash
make clean
```

1. Delete application binary [delete application build binary]

```bash
make clean_bin
```

## RDE support

```bash
Raise any issue or feature request using RDE Help
Join the Awesome Roost Community https://join.slack.com/t/roostai/shared_invite/zt-ea5mo10y-jDJgXiHn0RihSmucz0UZpw
```
