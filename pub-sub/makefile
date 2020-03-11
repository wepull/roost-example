SUBDIRS = producer consumer

all:
	make build
	make docker
	make deploy

build:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

.PHONY: docker
docker:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

deploy:
	kubectl apply -f kubenretes-manifests/consumer.yaml
	kubectl apply -f kubenretes-manifests/producer.yaml
	
clean:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done
	kubectl delete -f kubenretes-manifests/consumer.yaml --all --now=true
	kubectl delete -f kubenretes-manifests/producer.yaml --all --now=true
