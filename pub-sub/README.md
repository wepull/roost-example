# Pub-Sub Example

Pub-Sub application demonstrates producer and consumer based communciation with zbio as messaging platform.
Producer can create topics in zbio and start producing messages to those topics.
Messages produced by producers can be consumed by consumer by subscribing to those created topics.

There are 3 ways to interact with this Pub-Sub project

## 1. Default mode

Launch Roost Desktop Engine (RDE). Once Zettabytes K8s Environment  (ZKE) is up and running, open RKT Konsole (Roost Kubernetes Terminal) and run following commands -

> Right-click on `Makefile` and click `Run` for hasselfree deployments in ZKE

```bash
cd $GOPATH/src/github.com/roost-io/roost-example/pub-sub

# Build, dockerise and deploy into ZKE cluster
make

# View producer and consumer log (Use `Workload Analytics` available in RDE desktop for better insights on deployed application)
kubectl logs zbio-sample-producer --tail 300
kubectl logs zbio-sample-consumer --tail 300

# View zbio service logs
kubectl logs service/zbio-service --namespace zbio
```

Open RDE desktop's Workload Analytics to view application pods, services and logs

[RDE Workload Analytics image](Show pub-sub pod logs and workload view)

- > Default topics are created. i.e. `pub-sub-example-1` and `pub-sub-example-2`
- > Producer keeps producing sequencial messages to these default topis.
- > Consumer keeps consuming from default topics.

## 2. Interactive mode (producer and consumer can interact with user defined topics)

Once `producer` and `consumer` pods running (from default mode), execute following commands to interact with producer and consumer.

```bash
# Producer (single topic)
kubectl exec -it zbio-sample-producer -- producer  --interactive --topic=pub-sub-interactive-1 --message="Sent message from producer interactively"

# Producer (multiple topic)
kubectl exec -it zbio-sample-producer -- producer  --interactive --topic=pub-sub-interactive-1,pub-sub-interactive-1 --message="Sent message from producer interactively to multiple topics"

# Consumer (single topic)
kubectl exec -it zbio-sample-consumer -- consumer --interactive --topic=pub-sub-interactive-1

# Consumer (multiple topic)
kubectl exec -it zbio-sample-consumer -- consumer --interactive --topic=pub-sub-interactive-1,pub-sub-interactive-2
```

## 3. Prompt mode (producer and consumer prompt for user inputs)

Once `producer` and `consumer` pods running (from default mode), execute following commands to interact with producer and consumer. Press `crtl+c` and `n` to exit out of prompts.

```bash
# Producer
kubectl exec -it zbio-sample-producer -- producer  --prompt

# Consumer
kubectl exec -it zbio-sample-consumer -- consumer --prompt
```

- > Comma seperated topics are allowed.
- > If neither `--interactive` nor `--prompts` flags are provided, then producer and consumer run in **`default mode`** interacting with default topics i.e. `pub-sub-example-1` and `pub-sub-example-2`

## Cleaning (Always prefer to cleanup resources if not in use)

```bash
make clean
```

- > _Deletes producer and consumer binaries_
- > _Deletes deployed pub sub app resources from kubernetes_
