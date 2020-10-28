# Building docker images using kaniko as a Job in Kubernetes

## Pre-requisites before applying the job template
<ul>
<li>Make sure you have the source files and the Dockerfile for the image you are trying to build on the worker node where the job is scheduled to run</li>
<li>Mention the correct hostPath for the source files in the volume(named build-context) spec of the job template</li>
<li>Also apply the pv and pvc templates(for kaniko cache purposes) in this folder before applying the job yaml file</li>
</ul>