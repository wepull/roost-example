#MAKE fike for implementing example

#Build the docker image
build-image:
	docker build -t html_image:v1 .

#Run the docker image 
docker-run:
	docker run -p8080:8080 html_image:v1

#Deploy on the kubernetes cluster
deploy:
	kubectl apply -f my.yaml

#Clean and remove all the supported files
clean: 
	kubectl delete -f my.yaml
	docker rmi html_image:v1