all:
ifdef NODOCKER
	echo "DO NOT BUILD IMAGE"
else 
	docker build . -t gke-sample:v1
endif

.PHONY: deploy
deploy:
	kubectl apply -f deployment.yaml

.PHONY: clean
clean:
	kubectl delete -f deployment.yaml --all --now=true 