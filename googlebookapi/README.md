# GoogleBookAPI Example

Google book API application is build on top of [flogo](https://www.flogo.io/) application to demostrate use of zbio messaging platform in flow based application.
GoogleBookAPI application is integrated with zbio to send messages and persist them. These persisted messages are retained in zbio topics till rention period of the Topic. Messages can be further consumed by other component of the application or another microservices running. 

In this `GoogleBookAPI` example, zbio is used as logging activity in the flow to send every useful information in the form of messages to zbio. These messages are stored in the topics till topic's retention period. When application launch for first time, GoogleBookAPI application creates `googleBookAPI` topic to store messages. Every time request is made by the application to retrieve google book details, new activity logs are sent as new messages to ZBIO messaging platform in `googleBookAPI` topic.

## Let's run this application in RDE

Launch Roost Desktop Engine (RDE). Once Zettabytes Cluster Engine (ZKE) is up and running, open RKT Konsole (Roost Kubernetes Terminal) and run following commands -

> Right-click on `Makefile` and click `Run` for hasselfree deployments in ZKE

```bash
cd cd $GOPATH/src/github.com/roost-io/roost-example/googlebookapi

# Build, dockerise and deploy into ZKE cluster
make

# View application logs (Use `Workload Analytics` available in RDE desktop to get better insights on deployed application)
# -f :to keep streaming logs from application
kubectl logs -f service/googlebookapi

# zbio service logs
kubectl logs service/zbio-service --namespace zbio

# Deletes googlebookapi binaries and undeploy from ZKE
make clean
```

Open RDE desktop's Workload Analytics to view application pods, services and logs

[RDE Workload Analytics image](show_GoogleBookAPI_pod_logs_and_workload_view)

### How to access GoogleBookAPI application

* Open http://roost-master:30045/books/9788126568772
  * isbn_number: `9788126568772` ; URL: http://roost-master:30045/books/<isbn_number>
* ISBN stands for international standard book number , which is
    13 digit number uniquely identify all the books.
* ISBN can be found in internet. **ISBN:** `9781788999786`, **Book Name:** `Mastering Kubernetes`
* The digit should not have any special characters in between.

## Cleaning (Always prefer to cleanup resources if not in use)

```bash
# Deletes googlebookapi binaries and undeploy from ZKE
make clean
```
