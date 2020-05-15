package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ZB-io/zbio/client"
	"github.com/ZB-io/zbio/rpc/common"
)

var (
	zbCli               *client.Client
	zbioServiceEndpoint string = "zbio-service:50002"
	topicsIn                   = flag.String("topic", "", "Name of topic(s) to read messages from. Comma seperated list of topics are accepted. (*Required if --interactive)")
	interactiveIn              = flag.Bool("interactive", false, "(*Required) Start interactive session. Pass either --interactive or --prompt flag. (default: false)")
	promptIn                   = flag.Bool("prompt", false, "Prompts for inputs to enter in STDIN. No need to pass topic and message flags. Press crtl+c to exit. (ignored if --interactive flag is passed")
)

func init() {

	if zbsvc := os.Getenv("SERVICE_ADDRESS"); zbsvc != "" {
		zbioServiceEndpoint = zbsvc
	}

	cfg := client.Config{
		Name:            "pub-sub-consumer",
		ServiceEndPoint: zbioServiceEndpoint,
	}

	cli, err := client.New(cfg)

	if err != nil {
		log.Fatalf("Failed to get client. Error %v", err)
	} else {
		zbCli = cli
	}
}

func main() {
	flag.Parse()

	// If interactive session enabled, read flags or STDIN for inputs
	// crtl+c would enable to detach out of container
	if *interactiveIn {
		flagSubscriber()
	} else if *promptIn {
		promptSubscriber()
	} else {
		startConsumer()
	}
	// outputSubscription()
}

// startConsumer keep streaming messages from default topics.
func startConsumer() {
	log.Println("Consumer started.")

	// topics := []string{"test-topic-1", "test-topic-2"}
	topics := []string{"pub-sub-example-1", "pub-sub-example-2"}

	msgChanMap, errChan, err := zbCli.ReadMessages(zbCli.Name, zbCli.Name+"TestConsumerGroup", topics)

	if err != nil {
		log.Fatal("Read message failed " + err.Error())
	}

	go func() {
		for {
			select {
			case msg := <-msgChanMap[topics[0]]:
				log.Println("Message Received ", topics[0], string(msg))
			case msg := <-msgChanMap[topics[1]]:
				log.Println("Message Received ", topics[1], string(msg))
			case msg := <-errChan:
				log.Printf("Error reading messages from consumer channel. Error: %s", msg.Error())
			default:
				log.Println("Awaiting for response on consumer channel. Will sleep for 2 sec")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	select {}
}

// outputSubscription stream out topics data.
func outputSubscription(topics []string) {
	fmt.Println("Streaming response from topic(s)? Press ctrl+c exit and subscribe another topic")
	msgChanMap, errChan, err := zbCli.ReadMessages(zbCli.Name, zbCli.Name+"TestConsumerGroup", topics)
	if err != nil {
		log.Fatal("Read message failed " + err.Error())
	}

	go func() {
		for {
			select {
			case msg := <-msgChanMap[topics[0]]:
				log.Println("Message Received ", topics[0], string(msg))
			case msg := <-msgChanMap[topics[1]]:
				log.Println("Message Received ", topics[1], string(msg))
			case msg := <-errChan:
				log.Printf("Error reading messages from consumer channel. Error: %s", msg.Error())
			default:
				log.Println("Awaiting for response on consumer channel. Will sleep for 2 sec")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	select {}
}

func unsubscribe() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		promptSubscriber()
	}()
}

func sanitiseString(in string) string {
	return strings.Trim(in, "\n")
}

// subscriptionPrompt confirms if user wish to subscribe to topics
func subscriptionPrompt(reader *bufio.Reader) (string, error) {
	fmt.Println("\nWant to subscribe to topic? (y/n)")
	resp, err := reader.ReadString('\n')
	return sanitiseString(resp), err
}

// topicsPrompt accepts topic names from STDIN
func topicsPrompt(reader *bufio.Reader) ([]string, error) {
	var topics []string
	var topicsIn string
	var err error
	// Keep asking for topics if valid topics are not provided
	for {
		fmt.Println("\nEnter topic(s): \tComma seperated topic names are allowed")
		topicsIn, err = reader.ReadString('\n')
		if err != nil {
			return topics, err
		}

		topicsIn = sanitiseString(topicsIn)
		if topicsIn == "" {
			fmt.Println("Topic(s) not entered")
			continue
		}
		break
	}
	topics = strings.Split(sanitiseString(topicsIn), ",")
	return topics, err
}

// promptSubscriber interactively
func promptSubscriber() {
	reader := bufio.NewReader(os.Stdin)

	subscribe, err := subscriptionPrompt(reader)
	if err != nil {
		log.Fatal("Failed to read from STDIN")
	}

	switch subscribe {
	case "y":
		var topics []string
		topics, err := topicsPrompt(reader)
		if err != nil {
			log.Fatal("Unable to read user input from topicsPrompt")
			promptSubscriber()
		}

		go unsubscribe()

		// Check if Topics exists. If doesn't exists, keep asking for valid topic
		topicStatus, err := zbCli.DescribeTopic(topics)
		if err != nil {
			log.Printf("Error in Describing Topic %v\n", err)
			topics, err = topicsPrompt(reader)
			if err != nil {
				log.Fatal("Unable to read user input from topicsPrompt")
				promptSubscriber()
			}
		} else {
			for i, value := range topicStatus.GetTopics() {
				fmt.Printf("[%d.] These are the topic descriptions %v\n", i+1, value)
			}
		}

		// Do meaningful with topics subscription, when ReadMessage works fine
		outputSubscription(topics)

	case "n":
		os.Exit(0)

	default:
		fmt.Println("\nOnly y or n or quit is accepted.")
		for {
			promptSubscriber()
		}
	}
}

func flagSubscriber() {
	inputTopics := *topicsIn
	topics := strings.Split(inputTopics, ",")
	if len(topics) < 1 {
		log.Fatalf("Enter valid topic. You entered: %s\n", topics)
	}

	var validTopics []string
	for _, topic := range topics {
		ok := validate(topic)
		if !ok {
			log.Fatalf("Not a valid topic. You entered: %v\n", topic)
		}
	}

	topicStatus, err := zbCli.DescribeTopic(topics)
	if err != nil {
		log.Fatalf("Error describing to topic. Error: %v\n", err)
	}

	// Check if topics are found
	for topicName, response := range topicStatus.GetResponses() {
		if response.Code == common.ResponseStatus_OK {
			validTopics = append(validTopics, topicName)
			fmt.Printf("Topic: %s\tDescribe Response: %v\n", topicName, response.GetDetails())
		} else {
			log.Printf("Unable to describe topic: %s, Response Code: %v", topicName, response.Code)
		}
	}
	outputSubscription(validTopics)
}

func validate(topic string) bool {
	// Validate if not empty. Add more validations if any
	if topic != "" {
		return true
	}
	return false
}
