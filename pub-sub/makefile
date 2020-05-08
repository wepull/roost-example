SUBDIRS = producer consumer

all: clean build docker deploy

build:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

.PHONY: docker
docker:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

deploy:
	kubectl apply -f kubernetes-manifests/consumer.yaml
	kubectl apply -f kubernetes-manifests/producer.yaml
	
clean:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done
	kubectl delete -f kubernetes-manifests/consumer.yaml --now
	kubectl delete -f kubernetes-manifests/producer.yaml --now
