SUBDIRS = producer consumer

all: clean build docker deploy clean_bin

build:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

.PHONY: docker
docker:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

deploy:
	kubectl apply -f kubernetes-manifests/producer.yaml
	kubectl apply -f kubernetes-manifests/consumer.yaml
	
clean:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done
	kubectl delete -f kubernetes-manifests/consumer.yaml --now >/dev/null 2>&1 || true
	kubectl delete -f kubernetes-manifests/producer.yaml --now >/dev/null 2>&1 || true

.PHONY: clean_bin
clean_bin:
	rm -f producer/bin/*
	rm -f consumer/bin/*