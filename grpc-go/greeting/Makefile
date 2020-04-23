proto:
	protoc greetpb/greet.proto --go_out=plugins=grpc:.
server:
	go build greet_server/server.go
client:
	go build greet_client/client.go