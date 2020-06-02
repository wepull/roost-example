all: build dockerise

build:
	CGO_ENABLED=0 GOOS=linux go build -o fetcher main.go

dockerise: build
	@docker build -f Dockerfile -t fetcher:latest .
	${shell pwd}/../../push_image.sh fetcher:latest

deploy:
	@kubectl apply -f fetcher.yaml 

undeploy:
	@kubectl delete -f fetcher.yaml >/dev/null 2>&1 || true

clean:
	rm -f fetcher
