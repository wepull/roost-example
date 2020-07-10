# Health checking gRPC server on Kubernetes

This gRPC client-server app shows how to check health of gRPC servers on Kubernetes while using ZBIO messaging platform to persist grpc-client and grpc-server events.
Kubernetes health checks (liveness and readiness probes) detect unresponsive pods, mark them unhealthy, and cause these pods to be restarted or rescheduled.
However, Kubernetes does not support gRPC checks natively which means that the developers need to implement it.

gRPC has a standard health checking protocol that can be used from any language. 

## How to run this application

###### Using Roost Desktop Engine (RDE)

> Right-click on `Makefile` and click `Run` for hassle-free deployment to ZKE

 >What all is done by `make`?
  * Cleans existing deployment [kubernetes/deploy.yaml](kubernetes/deploy.yaml)
  * Removes the application binary [bin/grpcServer && bin/grpcClient](bin/grpcServer && bin/grpcClient)
  * Builds grpcClient and grpcServer by compiling the Golang code [src/main.go](src/main.go)
  * Builds an image using the app binary [serveri-grpci && client-grpc](server-grpci && client-grpc)
  * Deploys the image to ZKE Cluster [kubernetes/deploy.yaml](kubernetes/deploy.yaml)

###### Using RKT Konsole to run application

```bash
cd $GOPATH/src/github.com/roost-io/roost-example-latest/grpcExample
make
```

> User can also leverage multistage build apart from above build-deploy via `make`. (As base image contains protogenc and go, its size is larger ~450Mb, so download may take time)
>```bash
> cd $GOPATH/src/github.com/roost-io/roost-example-latest/grpcExample
> make multistage
> ```

## View application logs 
> Using `Workload Analytics` (RDE) for deployed application

> [RDE Workload Analytics image](show_GrpcExample_pod_logs_and_workload_view)

> Using RKT Konsole
```bash
kubectl get pods -n default # Get the grpc-deploy-* pod name
kubectl logs <grpc-deploy-XXX> -c client
kubectl logs <grpc-deploy-XXX> -c server
kubectl logs service/zbio-service --namespace zbio --tail 500

Additional options to logs can be given for streaming or getting last n lines
   * -f: to keep streaming logs from application
   * --tail <n>: to get the last n lines of output
```

## This sections provide explanations of each part of the application

### api

The server should export a service defined in the following proto for the health check. 
The following code which can be found [here](https://github.com/grpc/grpc/blob/v1.15.0/doc/health-checking.md) is added to the __api.proto__ file:

```proto
syntax = "proto3";

package grpc.health.v1;

message HealthCheckRequest {
  string service = 1;
}

message HealthCheckResponse {
  enum ServingStatus {
    UNKNOWN = 0;
    SERVING = 1;
    NOT_SERVING = 2;
  }
  ServingStatus status = 1;
}

service Health {
  rpc Check(HealthCheckRequest) returns (HealthCheckResponse);
}
```

Our server provides a `ProcessText` service which receives a message and the client name as `InputRequest` and sends a message and the server name as `OutputResponse`. We define the remote procedure call functions for this service in the __api.proto__ file:

```proto
// Defining the remote procedure call function that we want to be able to call on the data.
service ProcessText{
  rpc upper(InputRequest) returns (OutputResponse){}
}
// The serialized message that client sends.
message InputRequest{
  string text = 1;
  string clientName =2;
}
// What server responds to as a result of getting InputRequest.
message OutputResponse{
  string text = 1;
  string serverName =2;
}
```

### server-grpc

The server is given a specific name and it listens on a specific port for client requests. It provides the `Upper` service which takes the message received from the client and converts it to upper case and adds a smiley face emoji at the end of the string. Server sends the server name as a part of the `OutputResponse`.

```go
func (s server) Upper(c context.Context, req *api.InputRequest) (*api.OutputResponse, error) {
    x := happyUpper(req.GetText())
    log.Printf("➡️ Received message from client %v: %v ", req.GetClientName(), req.GetText())
    return &api.OutputResponse{ServerName: serverName, Text: x}, nil
}
```

The server has a dummy database in __db.go__ that has a readiness flag `isDatabaseReady` that is initially `false`. The dummy database waits up to a certain amount of time (e.g. 30 seconds) and then it changes the `isDatabaseReady` flag to `true`. In this application, the gRPC health check is implemented for the dummy database.

```go
func connectDB() error {
    sleepTime := 30
    log.Println("⏳ Connecting to the dummy database. This might take up to", sleepTime, "seconds")
    time.Sleep(time.Duration(sleepTime) * time.Second)
    log.Println("📣 Database is ready now!")
    isDatabaseReady = true
    return nil
}
```

If the `isDatabaseReady` flag is `true`, then `Check()` returns a `HealthCheckResponse_SERVING` status and if the flag is `false`, it returns a `HealthCheckResponse_NOT_SERVING` health check response. 

In general, there are four health check response serving status: `HealthCheckResponse_UNKNOWN`, `HealthCheckResponse_SERVING`, `HealthCheckResponse_NOT_SERVING` and `HealthCheckResponse_SERVICE_UNKNOWN`.

```go
func (h *Health) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
    log.Println("🏥 K8s is health checking")
    if isDatabaseReady == true {
      log.Printf("✅ Server's status is %s", grpc_health_v1.HealthCheckResponse_SERVING)
      return &grpc_health_v1.HealthCheckResponse{
        Status: grpc_health_v1.HealthCheckResponse_SERVING,
      }, nil
    } else if isDatabaseReady == false {
      log.Printf("🚫 Server's status is %s", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
      return &grpc_health_v1.HealthCheckResponse{
        Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
      }, nil
    } else {
      log.Printf("🚫 Server's status is %s", grpc_health_v1.HealthCheckResponse_UNKNOWN)
      return &grpc_health_v1.HealthCheckResponse{
        Status: grpc_health_v1.HealthCheckResponse_UNKNOWN,
      }, nil
    }
}
```

In the __startGrpcServer__ function, the following steps were taken to build and start the server:

1. Specify the port we want to use to listen for client requests using `ln, err := net.Listen("tcp", port)`. 
2. Create an instance of the gRPC server using `grpcServer := grpc.NewServer()`.
3. Register our service implementation with the gRPC server using `api.RegisterProcessTextServer(grpcServer, srv)`.
4. Register the health service using `grpc_health_v1.RegisterHealthServer(grpcServer, &Health{})`
5. Call Serve() on the server `err = grpcServer.Serve(ln)` with our port details to do a blocking wait until the process is killed or `Stop()` is called.

### server.Dockerfile

To use the gRPC standard health checking protocol, we need to include the compiled `grpc_health_probe` in the container image of the server.

```Dockerfile
# Use already downloaded grpc_health_probe
COPY grpc_health_probe-linux-amd64 /bin/grpc_health_probe
```

OR

```Dockerfile
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 
RUN chmod +x /bin/grpc_health_probe  
```

### client-grpc

The `main()` function in __main.go__ instantiates a client connection, on the TCP port the server is bound to.
You can specify an ip and port by using Command-Line Flag. For example:

```bash
./client -ip=127.0.0.1:3000
```

If Command-Line Flag is not used then the default ip and port will be used for the connection.

```go
ipPtr := flag.String("ip", "127.0.0.1:3000", "Description: ip address")
```

In __main.go__ a client has been created for the __ProcessText__ service: 

```go
client :=api.NewProcessTextClient(conn)
```

The client is given a name which is a random number.

```go
randomClientName := strconv.Itoa(seededRand.Intn(10000))
```

Integrates zbio messaging to persist messages in zbio topics. `zbutil.InitZBIO(zbConfig)` is responsible to create topics in zbio.

```go
zbConfig := zbutil.Config(randomClientName)
zbutil.InitZBIO(zbConfig)

message := zb.Message{
  TopicName:     zbutil.TopicName, // default topicName
  Data:          []byte(fmt.Sprintf("grpc-client starting from grpcExample. zbClientName: %s\n", randomClientName)),
  HintPartition: "",
}
zbutil.SendMessageToZBIO([]zb.Message{message})
```

The client sends the client name and a random message to the server periodically. It then receives the server name and the constructed message from the server and sleeps for 2 seconds. Messages sent from client to server and response received back from server are sent to zbio to persist them in zbio topics.

```go
for {

  var message zb.Message
  // Create a random string of length 10 to send to the server.
  randomMessage := randomString(10)

  // Create a context.
  ctx := context.Background()

  // Send a request.
  reqMessage := &api.InputRequest{
    Text:       randomMessage,
    ClientName: randomClientName}

  requestLog := fmt.Sprintf("⬅️ Client sent a message to server : %s", randomMessage)
  log.Println(requestLog)

  // Send requested message to zbio
  message = zb.Message{
    TopicName:     message.TopicName, // default topicName
    Data:          []byte(requestLog),
    HintPartition: "",
  }
  zbutil.SendMessageToZBIO([]zb.Message{message})

  // Receive response.
  resp, err := client.Upper(ctx, reqMessage)
  if err != nil {
    log.Printf("❌ Error doing upper : %v", err)
  }

  responseLog := fmt.Sprintf("➡️ Received Response from server %v : %s ", resp.GetServerName(), resp.Text)
  log.Printf(responseLog)

  // Send response message to zbio
  message = zb.Message{
    TopicName:     message.TopicName, // default topicName
    Data:          []byte(responseLog),
    HintPartition: "",
  }
  zbutil.SendMessageToZBIO([]zb.Message{message})

  // Sleep for 2 seconds before sending another message.
  time.Sleep(2 * time.Second)
}
```

## Zettabytes Kubernetes Engine (ZKE)

The __kubernetes/deploy.yaml__ defines the container and pod's spec to be deployed on ZKE cluster.
We use Kubernetes exec probes and define liveness and readiness probes for the gRPC server container.

For the `readinessProbe`, we use the command `/bin/grpc_health_probe`. 

```yaml
readinessProbe:
  exec:
    command: ["/bin/grpc_health_probe","-addr=:3000"]
  initialDelaySeconds: 10
  periodSeconds: 1
```

`initialDelaySeconds` indicates the number of seconds that kubelet should wait after the start of the container to performe the first health probe.

`periodSeconds` indicates how often, in seconds, the kubelet should perform the health probe.  

For the liveness probe, similarly we can use the gRPC health probe `/bin/grpc_health_probe`, or a command such as `cat /tmp/healthy` that if executes successflully it returns 0 and the container is considered alive.

Now you can use `kubectl get pods` to get a list of the pods and find the exact name of the pod which should start with __grpc-deploy__.

You can retrieve more information about each pod using `kubectl describe pod`. For example:

```bash
kubectl describe pod grpc-deploy-8f95984fd-qm2xd
```

This should show a message about the readiness probe failing `Readiness probe failed: service unhealthy (responded with "NOT_SERVING")`.

If you check the Readiness status of the server using `kubectl describe pod` before the database is ready, it should show `false`, otherwise if you check it after the database is ready, then it should show `true`.

To see the messages that are sent and received between client and server, you can use the following commands to print the logs of the client and server containers:

```bash
kubectl logs grpc-deploy-8f95984fd-qm2xd  -c client
kubectl logs grpc-deploy-8f95984fd-qm2xd  -c server
```

## References

1. https://kubernetes.io/blog/2018/10/01/health-checking-grpc-servers-on-kubernetes/
2. https://grpc.io/docs/tutorials/basic/go/
3. https://github.com/grpc/grpc/blob/v1.15.0/doc/health-checking.md
4. https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/
5. https://developers.google.com/protocol-buffers/docs/gotutorial
6. https://hub.docker.com/r/znly/protoc
7. https://github.com/grpc-ecosystem/grpc-health-probe

``` 
Raise any issue or feature request using RDE Help
Join the Awesome Roost Community https://join.slack.com/t/roostai/shared_invite/zt-ea5mo10y-jDJgXiHn0RihSmucz0UZpw
```
