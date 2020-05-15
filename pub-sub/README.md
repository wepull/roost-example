# Publisher Subscriber Example

This application demonstrates use of ZBIO messaging platform to implement a producer and consumer sample.

* _Producer_ can create topics and send messages related to those topics.
* Messages can be consumed by _Consumer_ by subscribing to those created topics.

There are 3 ways to interact with this Pub-Sub project

## 1. Default
###### Using Roost Desktop Engine (RDE)

> Right-click on `Makefile` and click `Run` for hassle-free deployment to ZKE

###### Using RKT Konsole to run application
```bash
cd $GOPATH/src/github.com/roost-io/roost-example/pub-sub
# Build, dockerise and deploy into ZKE cluster
make
```

- > Default topics are created. i.e. `pub-sub-example-1` and `pub-sub-example-2`
- > Producer keeps producing sequential messages to these default topics.
- > Consumer keeps consuming from default topics.

## 2. Interactive mode (producer and consumer can interact with user defined topics)

Once `producer` and `consumer` pods running (from default mode), execute following commands to interact with producer and consumer.

```bash
# Producer (single topic)
kubectl exec -it zbio-sample-producer -- producer  --interactive --topic=pub-sub-interactive-1 --message="Sent message from producer interactively"

# Producer (multiple topics)
kubectl exec -it zbio-sample-producer -- producer  --interactive --topic=pub-sub-interactive-1,pub-sub-interactive-1 --message="Sent message from producer interactively to multiple topics"

# Consumer (single topic)
kubectl exec -it zbio-sample-consumer -- consumer --interactive --topic=pub-sub-interactive-1

# Consumer (multiple topics)
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

- > Comma separated topics are allowed.
- > If neither `--interactive` nor `--prompts` flags are provided, then producer and consumer run in **`default mode`** interacting with default topics i.e. `pub-sub-example-1` and `pub-sub-example-2`


## View application logs 
> Using `Workload Analytics` (RDE)

> [RDE Workload Analytics image](Show pub-sub pod logs and workload view)

> Using RKT Konsole
```bash
kubectl logs -f zbio-sample-producer -n zbio
kubectl logs zbio-sample-producer --namespace zbio --tail 400
kubectl logs -f zbio-sample-consumer -n zbio
kubectl logs zbio-sample-consumer --namespace zbio --tail 400
kubectl logs service/zbio-service --namespace zbio
```
 - logs option
   * -f: to keep streaming logs from application
   * --tail <n>: to get the last n lines of output



``` 
Raise any issue or feature request using RDE 
Join the Awesome Roost Community https://join.slack.com/t/roostai/shared_invite/zt-ea5mo10y-jDJgXiHn0RihSmucz0UZpw
```
