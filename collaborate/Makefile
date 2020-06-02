SUBDIRS = fetcher retriever
all: build dockerise

build:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

dockerise: build
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

deploy:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

undeploy:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done

clean:
	for d in $(SUBDIRS); do echo "<==> Target: $@ for $$d <==>"; $(MAKE) -C $$d $@; echo "========="; done
