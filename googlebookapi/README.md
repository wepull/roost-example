# GoogleBookAPI Example

GoogleBookAPI is built on top of [flogo](https://www.flogo.io/) to demonstrate use of ZBIO messaging platform in a flow based application.

Upon launch for first time, the application creates a topic `googleBookAPI`. 
Every time a request is made by the user to retrieve any book detail, activity logs are generated and sent to ZBIO. These messages can be consumed by another component of the same application or a different micro-service in ZKE cluster.

## How to run this application

###### Using Roost Desktop Engine (RDE)

> Right-click on `Makefile` and click `Run` for hassle-free deployment to ZKE

 >What all is done by `make`?
  * Cleans existing deployment [src/googlebookapi.yaml](src/googlebookapi.yaml)
  * Removes the application binary [googlebookapi](googlebookapi)
  * Builds an executable by compiling the Go/Flogo code [src/main.go](src/main.go)
  * Builds an image using the app binary [src/Dockerfile](src/Dockerfile)
  * Deploys the image to ZKE Cluster [src/googlebookapi.yaml](src/googlebookapi.yaml)

###### Using RKT Konsole to run application

```bash 
cd $GOPATH/src/github.com/roost-io/roost-example/googlebookapi
make
```

## View application logs 
> Using `Workload Analytics` (RDE) for deployed application

> [RDE Workload Analytics image](show_GoogleBookAPI_pod_logs_and_workload_view)

> Using RKT Konsole
```bash
kubectl logs -f service/googlebookapi -n default
kubectl logs service/googlebookapi --namespace default --tail 400
   * -f: to keep streaming logs from application
   * --tail <n>: to get the last n lines of output
```
## How to access GoogleBookAPI
* ISBN stands for international standard book number.
* It is 13 digit number uniquely identify all the books.
* ISBN can be found in internet. 
 **ISBN:** `9781788999786`, **Book Name:** `Mastering Kubernetes`
* The digit should not have any special characters in between.

> Using any browser 
* For single-node cluster, try roost-worker instead of roost-controlplane
* Open http://roost-controlplane:30045/books/<isbn\>
>>Sample URL: http://roost-controlplane:30045/books/9788126568772

> Using RKT Konsole
  * For single-node cluster, try roost-worker instead of roost-controlplane
  * curl http://roost-controlplane:30045/books/<isbn\>
  >>Sample URL: http://roost-controlplane:30045/books/9788126568772

``` 
Raise any issue or feature request using RDE Help
Join the Awesome Roost Community https://join.slack.com/t/roostai/shared_invite/zt-ea5mo10y-jDJgXiHn0RihSmucz0UZpw
```
