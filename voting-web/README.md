Build voter/Dockerfile
Deploy voter/voter.yaml

And you can access roost-controlplane:30030 over a browser 
But this is just the UI


Build ballot/Dockerfile
Apply ballot/ballot.yaml to ZKE

And you can access roost-controlplane:30080 (GET)
But this is just the Ballot API

POST request can also be accessed at the same end-point

====
Run both microservices and you have a full-fledged voting app
