ZBCLIENT=zbclient
all: 
ifdef NODOCKER
	echo "DO NOT BUILD IMAGE"
else 
	docker build -f Dockerfile -t zbio/zbclient:v1 .
endif

clean:
	rm -rf ${ZBCLIENT}

