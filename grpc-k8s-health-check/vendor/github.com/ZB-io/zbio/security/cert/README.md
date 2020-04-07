
## Create cert
```
openssl genrsa -out cert/server.key 2048

openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650

~~Country Name (2 letter code) []:US
State or Province Name (full name) []:CA
Locality Name (eg, city) []:San Jose
Organization Name (eg, company) []:Zb.io
Organizational Unit Name (eg, section) []:Zbio
Common Name (eg, fully qualified host name) []:zb.io
Email Address []:sudhir@zb.io~~

openssl req -new -sha256 -key cert/server.key -out cert/server.csr

~~Country Name (2 letter code) []:US
State or Province Name (full name) []:CA
Locality Name (eg, city) []:San Jose
Organization Name (eg, company) []:Zb.io
Organizational Unit Name (eg, section) []:Zbio
Common Name (eg, fully qualified host name) []:zb.io
Email Address []:sudhir@zb.io

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:Zbio032020s~~

```

## local dns entry (pick your local IP) /etc/hosts
`10.1.16.52      zbio-service`

## standalone run
`go test -v . -run=TestCreateTopic -count=1`

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