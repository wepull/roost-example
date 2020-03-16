# ZBIO Knative integration

## Commands

### Install istio

```bash
export ISTIO_VERSION=1.3.6
curl -L https://git.io/getLatestIstio | sh -
# Add istioctl into PATH
istioctl verify-install

cd istio-${ISTIO_VERSION}

for i in install/kubernetes/helm/istio-init/files/crd*yaml; do kubectl apply -f $i; done
```

```bash
helm template --namespace=istio-system \
  --set prometheus.enabled=false \
  --set mixer.enabled=false \
  --set mixer.policy.enabled=false \
  --set mixer.telemetry.enabled=false \
  `# Pilot doesn't need a sidecar.` \
  --set pilot.sidecar=false \
  --set pilot.resources.requests.memory=128Mi \
  `# Disable galley (and things requiring galley).` \
  --set galley.enabled=false \
  --set global.useMCP=false \
  `# Disable security / policy.` \
  --set security.enabled=false \
  --set global.disablePolicyChecks=true \
  `# Disable sidecar injection.` \
  --set sidecarInjectorWebhook.enabled=false \
  --set global.proxy.autoInject=disabled \
  --set global.omitSidecarInjectorConfigMap=true \
  --set gateways.istio-ingressgateway.autoscaleMin=1 \
  --set gateways.istio-ingressgateway.autoscaleMax=2 \
  `# Set pilot trace sampling to 100%` \
  --set pilot.traceSampling=100 \
  --set global.mtls.auto=false \
  install/kubernetes/helm/istio \
  > ./istio-lean.yaml

  kubectl apply -f istio-lean.yaml
```

### Install knative

```bash

kubectl apply --filename https://github.com/knative/serving/releases/download/v0.13.0/serving.yaml \
--filename https://github.com/knative/eventing/releases/download/v0.13.0/eventing.yaml \
--filename https://github.com/knative/serving/releases/download/v0.13.0/monitoring.yaml

```

### Attach knative brokers into namespace

```bash

```

```bash
kubectl get broker -n zbio
```

### Curl Producer

```bash
kubectl --namespace zbio apply --filename - << END
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: curl
  name: curl
spec:
  containers:
    # This could be any image that we can SSH into and has curl.
  - image: radial/busyboxplus:curl
    imagePullPolicy: IfNotPresent
    name: curl
    resources: {}
    stdin: true
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    tty: true
END
```

```bash
kubectl run --generator=run-pod/v1 curl --image=radial/busyboxplus:curl -it -n zbio
```

```bash
curl -v "http://default-broker.zbio.svc.cluster.local" \
-X POST \
-H "Ce-Id: unique-request-id-1" \
-H "Ce-Specversion: 0.3" \
-H "Ce-Type: newMessage" \
-H "Ce-Source: broker" \
-H "Content-Type: application/json" \
-d '{"msg":"Sending message to knative default broker - 1"}'
```

### Watch received events into app

```bash
kubectl --namespace zbio logs -l app=broker-1-consumer --tail=100
```
