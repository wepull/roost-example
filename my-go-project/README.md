#

## Commands

- Build application with `golang:latest` as base image.

```bash
    cd $GOPATH/src/github.com/roost-io/roost-example/my-go-project

    # build image with golang as base image
    docker build -t html:golang .

    # multistage build with alpine as base image
    docker build -f alpine.Dockerfile -t html:alpine .

    # multistage build with alpine as base image
    docker build -f scratch.Dockerfile -t html:scratch .

    # compare all image sizes
    docker images | grep html
```
