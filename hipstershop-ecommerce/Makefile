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

.PHONY: all
all: clean dockerise deploy

.PHONY: dockerise
dockerise:
ifeq ($(HOSTNAME), roost-utility)
	for service in ${MICROSERVICES}; do \
		docker ps >/dev/null 2>&1 && docker build -f src/$${service}/Dockerfile -t $${service}:${IMAGE_TAG} ./src/$${service} ; \
		@sh ${shell pwd}/../push_image.sh "$${service}:${IMAGE_TAG}" ; \
	done
else
	for service in ${MICROSERVICES}; do \
		echo "Building $${service}"; \
		docker build -f src/$${service}/Dockerfile -t $${service}:${IMAGE_TAG} ./src/$${service} ; \
	done
endif

.PHONY: deploy
deploy:
	kubectl apply -f kubernetes-manifests/zke-deployment.yaml

.PHONY: clean
clean:
	kubectl delete -f kubernetes-manifests/zke-deployment.yaml || true