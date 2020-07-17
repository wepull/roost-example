# ./kubernetes-manifests

:warning: Kubernetes manifests provided in this directory are not directly
deployable to a cluster. They are meant to be used with `skaffold` command to
insert the correct `image:` tags.

Use the manifests in [/release](/release) directory which are configured with
pre-built public images.

## To deploy in ZKE cluster

Use [/kubernetes-manifests/zke-deployment.yaml](/kubernetes-manifests)
Application can open in browser over http://roost-controlplane:30046

>_Note:_ Applicaiton is exposed over nodePort 30046
