# Run Client Test

## Build
```
make clean
make build
```

## local dns entry (pick your local IP) /etc/hosts
`10.1.16.52      zbio-service`

## standalone run
`SERVER_CRT="$GOPATH/src/github.com/ZB-io/zbio/security/cert/server.crt" bash -c 'go test -v . -run=TestCreateTopic -count=1'"`

## docker
```
make clean
make build-linux

#pick the latest abbrevated commit log
git log --abbrev-commit

#example commit id: a3876b7

docker build -t zbio/zbio:client-a3876b7 .
docker push zbio/zbio:client-a3876b7

make containerise

```


## Kubernetes
```
make clean
make build-linux

#pick the latest abbrevated commit log
git log --abbrev-commit

#example commit id: a3876b7

docker build -t zbio/zbio:client-a3876b7 .
docker push zbio/zbio:client-a3876b7

#open kubernetes-manifests/zbio-client.yaml
#change broker and service images 
#example: image: zbio/zbio:client-a3876b7

docker pull zbio/zbio:client-a3876b7

kubectl apply -f ../kubernetes-manifests/zbio-client.yaml 

kubectl get all 
kubectl get deployment
kubectl get pod 
kubectl describe pod/zbio-client

kubectl logs pod/zbio-client

```