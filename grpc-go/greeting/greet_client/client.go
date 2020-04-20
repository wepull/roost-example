package main 

import (
  "fmt"
  "log"
  
  greetpb "github.com/roost-io/examples/grpc-go/greeting/greetpb"
  
    "google.golang.org/grpc"
)

func main() {
  fmt.Println("Hellow I am a client")
  conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure)
  if err != nil {
    log.Fatalf("Could not connect: %v",err)
    
    defer cc.close() 
    
    c := greetpb.NewGreetServiceClient(cc)
  }
}