package main

import (
  "net"
  "log"
  
  pb "github.com/roost-io/roost-example/grpc-go/greeting/greetpb"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

type Greeter struct{}

func (s *Greeter) SayGreeting(context context.Context, in *pb.GreetRequest)(*pb.GreetResponse,error) {
  return &pb.GreetResponse{Message: "Hello" + in.Name},nil
}

func main() {
  lis, err := net.Listen("tcp","0.0.0.0:50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterGreetServiceServer(s,&Greeter{})
  
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}