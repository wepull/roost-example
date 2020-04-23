package main 

import (
  "context"
  "log"

  pb "github.com/roost-io/roost-example/grpc-go/greeting/greetpb"
  
    "google.golang.org/grpc"
)

func main() {
  conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect: %v",err)
  }
    c := pb.NewGreetServiceClient(conn)
    greeting, err := c.SayGreeting(context.Background(),&pb.GreetRequest{})
    log.Print(greeting)
  
  
}