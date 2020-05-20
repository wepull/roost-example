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

	"github.com/ZB-io/zbio/client"
)

var (
	zbCli               *client.Client
	zbioServiceEndpoint string = "zbio-service.zbio:50002"
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
}

// startConsumer keep streaming messages from default topics.
func startConsumer() {
	topics := []string{"pub-sub-example-1", "pub-sub-example-2"}

	msgChanMap, err := zbCli.ReadMessages("pub-sub-testClient", "pub-sub-TestClientGroup", topics)

	if err != nil {
		log.Fatalf("Read message failed " + err.Error())
	}

	for topic, msgChan := range msgChanMap {
		go printConsumedMessages(topic, msgChan)
	}
	select {}
}

func printConsumedMessages(topic string, msgChan chan []byte) {
	for {
		msg, ok := <-msgChan
		if !ok {
			log.Printf("Read channel closed for topic: %s\n", topic)
			break
		} else {
			log.Printf("Message from topic: %s\t is: %v\n", topic, string(msg))
		}
	}
}

// outputSubscription stream out topics data.
func outputSubscription(topics []string) {
	log.Println("Streaming response from topic(s)? Press ctrl+c exit and subscribe another topic")
	msgChanMap, err := zbCli.ReadMessages("pub-sub-interactive-testClient", "pub-sub-interactive-TestClientGroup", topics)

	if err != nil {
		log.Fatalf("Read message failed " + err.Error())
	}

	for topic, msgChan := range msgChanMap {
		go printConsumedMessages(topic, msgChan)
	}
	select {}
}

func unsubscribe() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
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
	go unsubscribe()
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

	for _, topic := range topicStatus.GetTopics() {
		validTopics = append(validTopics, topic.GetName())
	}

	log.Printf("Valid topics are %v", validTopics)
	outputSubscription(validTopics)
}

func validate(topic string) bool {
	// Validate if not empty. Add more validations if any
	if topic != "" {
		return true
	}
	return false
}
