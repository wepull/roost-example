package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ZB-io/zbio/client"
)

var (
	zbCli               *client.Client
	topicName           string = "pub-sub-example"
	zbioServiceEndpoint string = "zbio-service:50002"
)

func init() {
	if zbsvc := os.Getenv("SERVICE_ADDRESS"); zbsvc != "" {
		zbioServiceEndpoint = zbsvc
	}
	cfg := client.Config{
		Name:            "TestProducer",
		ServiceEndPoint: zbioServiceEndpoint,
	}
	cli, err := client.New(cfg)

	if err != nil {
		fmt.Println("Failed to get client, err=", err)
	} else {
		zbCli = cli
	}
}

func createTopics() {
	var topics = []string{"pub-sub-example-1", "pub-sub-example-2"}
	ok, err := zbCli.CreateTopics(topics, "", 3, 1, 100000)
	if !ok {
		log.Fatalf("Unable to create Topics due to error: %v", err)
	}
}

/*
TODO:
	1. No need to sleep for a sec.
*/
func sendNewMessage() {
	messages := []client.Message{
		client.Message{
			TopicName:     "pub-sub-example-1",
			HintPartition: "",
			Data:          []byte("Message number 0"),
		},
		client.Message{
			TopicName:     "pub-sub-example-2",
			HintPartition: "",
			Data:          []byte("Message number 0"),
		},
	}

	// Send 1000 messages
	for i := 0; i < 9999; i++ {
		messages[0].Data = []byte(fmt.Sprintf("Message number %d ", i))
		messages[1].Data = []byte(fmt.Sprintf("Topic Message count %d ", i))

		response, err := zbCli.NewMessage(messages)
		if err != nil {
			log.Fatalf("NewMessage failed with error: %v", err)
		}
		log.Printf("NewMessage Response: %v", response)

		time.Sleep(1 * time.Second)
	}
}

func startProducer() {
	createTopics()
	sendNewMessage()
}

func main() {
	startProducer()
}
