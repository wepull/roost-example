EMAIL_SERVICE_IMG=emailservice
PRODUCT_CATALOG_SERVICE_IMG=productcatalogservice
RECOMMENDATION_SERVICE_IMG=recommendationservice
SHIPPING_SERVICE_IMG=shippingservice
CHECKOUTSERVICE_IMG=checkoutservice
PAYMENT_SERVICE_IMG=paymentservice
CURRENCY_SERVICE_IMG=currencyservice
CARTSERVICE_IMG=cartservice
FRONTEND_IMG=frontend
LOAD_GENERATOR_IMG=loadgenerator
ADDSERVICE_IMG=adservice

IMAGE_TAG=v1

GRPC_HEALTH_PROBE=grpc_health_probe-linux-amd64
GRPC_HEALTH_PROBE_PATH=../tool/${GRPC_HEALTH_PROBE}

HOSTNAME := $(shell hostname -f)

GO_SERVICES = ${PRODUCT_CATALOG_SERVICE_IMG} ${SHIPPING_SERVICE_IMG} ${CHECKOUTSERVICE_IMG} ${FRONTEND_IMG}
MICROSERVICES = ${EMAIL_SERVICE_IMG} \
				${PRODUCT_CATALOG_SERVICE_IMG} \
				${RECOMMENDATION_SERVICE_IMG}	\
				${SHIPPING_SERVICE_IMG}	\
				${CHECKOUTSERVICE_IMG}	\
				${PAYMENT_SERVICE_IMG}	\
				${CURRENCY_SERVICE_IMG}	\
				${CARTSERVICE_IMG}	\
				${FRONTEND_IMG}	\
				${LOAD_GENERATOR_IMG}	\
				${ADDSERVICE_IMG}

SERVICE_GRPC_BIN = $(foreach microservice, ${MICROSERVICES}, "src/${microservice}"/${GRPC_HEALTH_PROBE})
GO_SERVICE_BIN = $(foreach go_services, ${GO_SERVICES}, "src/${go_services}"/${go_services})

.PHONY: all
all: clean prerequisites build dockerise deploy
	$(MAKE) clean_bin

.PHONY: prerequisites
prerequisites:
	for service in ${MICROSERVICES}; do \
		if ( [ $${service} != ${FRONTEND_IMG} ] && [ $${service} != ${LOAD_GENERATOR_IMG} ] ); then \
			cp ${GRPC_HEALTH_PROBE_PATH} src/$${service}; \
		fi; \
	done

.PHONY: build
build:
	for d in $(GO_SERVICES); do echo "<==> Target $@ for src/$${d} <==>"; $(MAKE) -C src/$${d} $@; echo "========="; done

.PHONY: dockerise
dockerise:
	for service in ${MICROSERVICES}; do \
		docker build -f src/$${service}/Dockerfile -t $${service}:${IMAGE_TAG} ./src/$${service} ; \
	done

.PHONY: deploy
deploy:
	kubectl apply -f kubernetes-manifests/zke-deployment.yaml

.PHONY: clean
clean: clean_bin 
	kubectl delete -f kubernetes-manifests/zke-deployment.yaml --now >/dev/null 2>&1 || true

.SILENT:
clean_bin:
	@(for grpc_bin in ${SERVICE_GRPC_BIN}; do \
		rm -f $${grpc_bin}; done;	\
	)
	@(for go_service_bin in ${GO_SERVICE_BIN}; do \
		if [ -f $${go_service_bin} ] ; then \
			rm -f $${go_service_bin};	\
		fi;	\
	done; \
	)
