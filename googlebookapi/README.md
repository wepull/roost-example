# GoogleBookAPI Example

Google book API application is built on top of [flogo](https://www.flogo.io/) application to demonstrate use of ZBIO messaging platform in flow based application.
GoogleBookAPI application is integrated with ZBIO to send messages and persist them. These messages are retained in ZBIO topics till retention period of the Topic.
Messages can be further consumed by another component of the application or a different micro-service.

In this example, ZBIO is used as logging activity in the flow to send every useful information in the form of messages to ZBIO. 
Upon launch for first time, GoogleBookAPI application creates topic `googleBookAPI`. 
Every time a request is made by the user to retrieve any book detail, activity logs are sent as messages to ZBIO messaging platform under `googleBookAPI` topic.

## How to run this application

Using Roost Desktop Engine (RDE).

> Right-click on `Makefile` and click `Run` for hassle-free deployments in ZKE

 What all is done by `make`?
  * Cleans existing deployment (if any)
  * Removes executable
  * Builds an executable by compiling the Go/Flogo code
  * Containerize the executable
  * Deploy the image to ZKE Cluster

> Using RKT Konsole to run application
```bash 
cd $GOPATH/src/github.com/roost-io/roost-example/googlebookapi
make

## View application logs 
> Using `Workload Analytics` (RDE) for deployed application

> [RDE Workload Analytics image](show_GoogleBookAPI_pod_logs_and_workload_view)

> Using RKT Konsole
```bash
kubectl logs -f service/googlebookapi -n zbio
kubectl logs service/googlebookapi --namespace zbio --tail 400
   * -f: to keep streaming logs from application
   * --tail <n>: to get the last n lines of output


## How to access GoogleBookAPI application
* ISBN stands for international standard book number.
* It is 13 digit number uniquely identify all the books.
* ISBN can be found in internet. 
 **ISBN:** `9781788999786`, **Book Name:** `Mastering Kubernetes`
* The digit should not have any special characters in between.

> Using any browser
* Open http://roost-master:30045/books/<isbn\>
>>sample URL: http://roost-master:30045/books/9788126568772

> Using RKT Konsole
  * curl http://roost-master:30045/books/<isbn\>
  >>sample URL: http://roost-master:30045/books/9788126568772

``` 
Raise any issue or feature request using RDE 
Join the Awesome Roost Community
