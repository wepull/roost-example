package main

import (
	"fmt"
	"log"

	"github.com/ZB-io/zbio/client"
)

var (
	zbCli *client.Client2020-03-05T10:21:32.623+0530	info	client/client.go:323	Requested NewMessage data: [topic:"test-topic-1" data:"Message number 56 " ]

)

func init() {

	cfg := client.Config{

		Name:            "TestProducer",
		ServiceEndPoint: "zbio-service:50002",
	}
	cli, err := client.New(cfg)

	if err != nil {
		fmt.Println("Failed to get client, err=", err)
	} else {
		zbCli = cli
	}
}

func main() {

	startConsumer()

}

func startConsumer() {

	topics := []string{"test-topic-1", "test-topic-2"}

	msgChanMap, err := zbCli.ReadMessages("testClient", "TestClientGroup",
		topics)

	if err != nil {
		log.Fatal("Read message failed " + err.Error())
	}

	go func() {
		for {
			select {
			case msg := <-msgChanMap[topics[0]]:
				fmt.Println("Message Received ", topics[0], string(msg))

			case msg := <-msgChanMap[topics[1]]:
				fmt.Println("Message Received ", topics[1], string(msg))

			}
		}
	}()

	select {}

}
