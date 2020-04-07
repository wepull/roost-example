# Health checking gRPC server on Kubernetes

This gRPC client server application was implemented for the purpose of showing how to do the health check of gRPC servers on Kubernetes.

Kubernetes health checks (liveness and readiness probes) detect unresponsive pods, mark them unhealthy, and cause these pods to be restarted or rescheduled.

Kubernetes does not support gRPC health checks natively which means that the developers should implement it when they deploy to Kubernetes. 

gRPC has a standard health checking protocol that can be used from any language. In this example, we have implemented this standard health checking protocol in our gRPC app, and invoked the `Check()` method to determine the server's status.

The next sections provide explanations of each part of the application.

## Commands to execute the Health checking gRPC on Kubernetes

For removing the binaries and dependencies on the docker daemon
	In Server part : make clean-server
	In Client part : make clean-client

For building the docker images 
	In server part : make build-server-img
	In client part : make build-client-img

For executing the docker images 
	In server part : make run-server
	In client part : make run-client
Note: Client and server codes need to be executed on the different tabs

For deploying on the kubernetes cluster 
	make deploy
