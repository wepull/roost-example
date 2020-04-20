package main

import (
  greetpb "github.com/roost-io/examples/grpc-go/greeting/greetpb"
  "net"
  "fmt"
  "log"
  
  "google.golang.org/grpc"
)

type server struct{}

func main() {
  fmt.Println("Hello World")
  lis, err := net.Listen("tcp","0.0.0.0:50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }
  s := grpc.NewServer()
  greetpb.RegisterGreetServiceServer(s,&server)
  
  if err := s.Server(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}