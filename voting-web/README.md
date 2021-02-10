Build voter/Dockerfile
Deploy voter/voter.yaml

And you can access roost-controlplane:30030 over a browser 
* For single-node cluster, try roost-worker instead of roost-controlplane
But this is just the UI


Build ballot/Dockerfile
Apply ballot/ballot.yaml to ZKE

And you can access roost-controlplane:30080 (GET)
* For single-node cluster, try roost-worker instead of roost-controlplane
But this is just the Ballot API

POST request can also be accessed at the same end-point

====
Run both microservices and you have a full-fledged voting app


## Service dependency test suite

Build and deploy service-test-suite in roost cluster.
Upload service-dependency.json from Observability -> Service Fitness in Roost Desktop
On building ballot image from within Roost Desktop, service filness events can be seen from event viewer.