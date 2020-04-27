SUBDIRS = service broker client scratchpad/zbio

all:
	make build-linux
	make image

build:
	for d in $(SUBDIRS); do echo "<==> Target $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

build-linux:
	for d in $(SUBDIRS); do echo "<==> Target $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

image:  
	# make -C images
	for d in $(SUBDIRS); do echo "<==> Target $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

clean:
	for d in $(SUBDIRS); do echo "<==> Target $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done
	find rpc -name '*.pb.go' -exec rm -f {} \;
	rm -r rpc/broker rpc/serviceint
	rm -rf zbimages.tar ./zbimages

save:
	rm -rf zbimages.tar ./zbimages
	mkdir -p ./zbimages
	docker save zbio/zbbroker:v1 > ./zbimages/zbbroker.tar
	docker save zbio/zbsvc:v1 > ./zbimages/zbservice.tar
	tar -cvf zbimages.tar ./zbimages

