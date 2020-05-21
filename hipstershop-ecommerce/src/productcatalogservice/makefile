MICROSERVICE_NAME=productcatalogservice

build:
	$(shell CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${MICROSERVICE_NAME} -i -mod vendor -modfile ../../go.mod -gcflags='-N -l' github.com/roost-io/roost-example/hipstershop-ecommerce/src/${MICROSERVICE_NAME})
	@(if [ -f ${MICROSERVICE_NAME} ]; then \
		echo "build successful";	\
	else \
		echo "Error: build not found" ;	\
	fi; \
	)