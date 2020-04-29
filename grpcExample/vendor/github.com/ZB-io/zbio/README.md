# Zbio
**POD Native Messaging with built-in Observability**

## Objectives
- Implement Distributed/**POD native messaging** system in GO
- **Cloud-native messaging**
- Kubernetes is designed to run stateless/lightweight workloads, need a native solution for messaging with lightweight containers
- Messaging will provide integrated **instrumentation/observability**, no need of sidecars, no need of separate messaging solution

## Features, beta
- Native Distributed Messaging - Stateless
  - Publish/Subscribe
  - Queueing
  - Request/Reply
- Registration, Discovery, Replication
- Configuration, management, state - API Server/config maps
- In-memory/cache/hot data
- Auto Instrumentation - via messaging client API/REST
- Realtime Telemetry - Prometheus, Grafana
- Intelligent Traffic routing (ML) - latency, traffic, errors, saturation
- Webhooks for sidecar proxy
- Go Client
- ZB Dashboard mostly for messaging
- ZB Explorer - basic

## Features, Phase 2
- CLI/REST Client
- Streaming
- Webhooks for service mesh
- ZB Explorer - advanced

## TODO
- [x] Objectives
- [x] Architecture
    - [ ] Consensus - need of etcd/zookeeper - no?
        - [ ] Raft
        - [ ] BFT Raft
    - [ ] Open Telemetry integration
    - [ ] Influx/prometheus for streaming/metrics
    - [ ] in-memory DB: mem cache DB
    - [ ] Linkedlist
    - [ ] Consumer groups, no topic (sync), partition, offset/checkpoint
- [ ] Code Structure

## Project/Code Layout

```
├── broker          broker system
├── cli             all CLI/Yaml
├── common          all CLI/Yaml
 ├── proto 
├── config          all config/
├── event         
├── message       
│   ├── protocol   
│   └── message    
├── metadata   
├── metrics    
├── observe         observability   
├── policy      
├── script     
├── security
    ├── proto      
├── sidercar-injector
├── state           memdb or etcd or in-memory
├── stream          streaming
└── topic           queue via partition, groups, sync via no topic names.      
└── util         
```

## Building

### Local

1. Clone zbio

    ```
    $ go get github.com/zb-io/zbio
    ```

2. Build zbio

    ```
    $ cd $GOPATH/src/github.com/zb-io/zbio
    $ make build
    ```

    (make sure that `$GOPATH/bin` is in your `PATH`)

### Docker

`docker build -t ... .`

---
## Reading, Frameworks
- Sample code bases in GO
    -  https://github.com/travisjeffery/jocko
    -  https://github.com/FireEater64/gamq
- http://studentnet.cs.manchester.ac.uk/resources/library/3rd-year-projects/2016/george.vanburgh.pdf
- https://www.mstakx.com/wp-content/uploads/2018/09/A-Practical-Observability-Primer-1.pdf
- http://www.mammatustech.com/kafka-architecture
- https://github.com/operator-framework/awesome-operators
- [Kafka Architecture] https://sookocheff.com/post/kafka/kafka-in-a-nutshell/
- [Pod native messaging, zb.io] (https://zb.io)
- [Best Microservice Frameworks (Go Kit?)] https://medium.com/seek-blog/microservices-in-go-2fc1570f6800
- https://nordicapis.com/7-frameworks-to-build-a-rest-api-in-go/
- https://dzone.com/articles/golang-guide-a-list-of-top-golang-frameworks-ides
- https://oxozle.com/awetop/avelino-awesome-go/
- https://streaml.io/blog/pulsar-segment-based-architecture
- https://www.confluent.io/blog/apache-kafka-for-service-architectures/

## Coding guidelines/best practices
- https://github.com/bahlo/go-styleguide
- https://github.com/smallnest/go-best-practices
- https://medium.com/@teivah/good-code-vs-bad-code-in-golang-84cb3c5da49d

## Need to cover
- Distributed Messaging
- In-memory Layer
- Instrumentation
- Service Discovery (eventual consistency, distributed caching)
- Load Balancing (least request, consistent hashing, zone/latency aware)
- Communication Resiliency (retries, timeouts, circuit-breaking, rate limiting)
- Security (end-to-end encryption, authorization policies, service-to-service ACLs)
- Observability (Layer 7 metrics, tracing, alerting)
- Routing Control (intelligent traffic shifting and mirroring)
- APIs
- Fault Injection (adding a timeout or error to test resiliency)
- Alerting

## Licensing
- TBD
