all:  dockerise undeploy deploy
	${MAKE} clean

build:
	CGO_ENABLED=0 GOOS=linux go build -o fetcher main.go

dockerise:
	@docker build -f Dockerfile -t fetcher:latest .
	${shell pwd}/../../push_image.sh fetcher:latest

deploy:
	@kubectl apply -f fetcher.yaml 

undeploy:
	@kubectl delete -f fetcher.yaml >/dev/null 2>&1 || true

clean:
	- rm -f fetcher
