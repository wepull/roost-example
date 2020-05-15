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
)

var (
	zbCli               *client.Client
	topicName           string = "pub-sub-example"
	zbioServiceEndpoint string = "zbio-service:50002"
	topicsIn                   = flag.String("topic", "", "Name of topic(s) to send messages on. Comma seperated list of topics are accepted. (*Required if --interactive)")
	messageIn                  = flag.String("message", "", "Message to send on topic(s). (*Required if --interactive) ")
	interactiveIn              = flag.Bool("interactive", false, "(*Required) Start interactive session. Pass either --interactive or --prompt flag. (default: false)")
	promptIn                   = flag.Bool("prompt", false, "Prompts for inputs to enter in STDIN. No need to pass topic and message flags. Press crtl+c to exit. (ignored if --interactive flag is passed")
)

func init() {
	if zbsvc := os.Getenv("SERVICE_ADDRESS"); zbsvc != "" {
		zbioServiceEndpoint = zbsvc
	}
	cfg := client.Config{
		Name:            "pub-sub-producer",
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
		shutDown()
		flagProducer()
	} else if *promptIn {
		shutDown()
		for {
			interactiveProducer()
		}
	} else {
		startProducer()
	}
}

// flagProducer accepts topicNames and messages from user as flag parameter to produce messages
func flagProducer() {
	var topics []string
	if *topicsIn == "" {
		fmt.Println("Topic is empty. Exiting...")
		os.Exit(1)
	} else {
		topics = strings.Split(sanitiseString(*topicsIn), ",")
	}

	if *messageIn == "" {
		fmt.Println("Message is empty. Exiting...")
		os.Exit(1)
	}

	// Create topic before sending message to those topics
	topicStatus := createTopics(topics)
	fmt.Printf("Create topic status?\tStatus: %v", topicStatus)

	// Send message to topic(s)
	response := sendMessages(topics, *messageIn)

	for topicName, status := range response {
		tName := strings.Split(topicName, ":")
		log.Printf("Producer's send message status?\tTopicName: %s\tStatus: %s\n", tName[0], status)
	}
	os.Exit(0)
}

// startProducer creates 2 topics and keep sending sequential messages to those topics
func startProducer() {
	// Create Topic
	var topics = []string{"pub-sub-example-1", "pub-sub-example-2"}
	ok, err := zbCli.CreateTopics(topics, "", 3, 1, 100000)
	if !ok {
		log.Fatalf("Unable to create Topics due to error: %v", err)
	}

	// Send Messages
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
			log.Fatalf("NewMessage failed with error: %v\n", err)
		}
		for resp, status := range response {
			// msgInfo := strings.Split(resp, ",") // topicName:partition:messageIndex
			log.Printf("messages sent to zbio status? TopicName: %s\t Status: %s", resp, status)
		}

		time.Sleep(1 * time.Second)
	}
	// Let producer running if user wants to send produce messages interactively
	time.Sleep(3 * time.Hour)
}

// newMessagePrompt confirms if user wish to send messages to topics
func newMessagePrompt(reader *bufio.Reader) (string, error) {
	fmt.Println("\nWant to send message to topic? (y/n)")
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
		fmt.Println("\nEnter topic(s) : (Comma seperated topic names are allowed)")
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

// topicMessagePrompt takes message to send to topic from STDIN
func topicMessagePrompt(reader *bufio.Reader) (string, error) {
	var message string
	var err error
	fmt.Println("\nEnter Message:\t(Message would be sent to above topics)")
	message, err = reader.ReadString('\n')
	if err != nil {
		return message, err
	}
	message = sanitiseString(message)
	return message, err
}

// interactiveProducer accepts inputs from STDIN to send messages to topics
func interactiveProducer() {
	var err error
	reader := bufio.NewReader(os.Stdin)

	isNewMessage, err := newMessagePrompt(reader)
	if err != nil {
		log.Fatal("Unable to read user input from newMessagePrompt")
	}
	switch isNewMessage {
	case "y":
		var topics []string
		var message string

		topics, err = topicsPrompt(reader)
		if err != nil {
			log.Fatal("Unable to read user input from topicsPrompt")
		}

		message, err = topicMessagePrompt(reader)
		if err != nil {
			log.Fatal("Unable to read user input for message to send to topic")
		}

		// Create topic before sending message to those topics
		topicStatus := createTopics(topics)
		fmt.Printf("Create topic status?\tStatus: %v", topicStatus)

		// Send message to topic(s)
		response := sendMessages(topics, message)
		for topicName, status := range response {
			tName := strings.Split(topicName, ":")
			log.Printf("Producer's send message status?\tTopicName: %s\tStatus: %s\n", tName[0], status)
		}
		fmt.Printf("\n===========================\n\n")

	case "n":
		// fmt.Printf("\nEntered: %v", "I would remain listening unless program closes..")
		os.Exit(0)

	default:
		fmt.Println("\nOnly y or n is accepted.")
	}
}

// createTopics creates topics with default configurations
func createTopics(topics []string) bool {
	ok, err := zbCli.CreateTopics(topics, "", 3, 1, 100000)
	if !ok {
		log.Fatalf("Unable to create Topics due to error: %v", err)
	}
	return ok
}

// sendMessages to topics.
func sendMessages(topics []string, message string) map[string]string {
	var topicMessges []client.Message
	for _, topic := range topics {
		messages := client.Message{
			TopicName:     topic,
			HintPartition: "",
			Data:          []byte(message),
		}
		topicMessges = append(topicMessges, messages)
	}

	fmt.Printf("\nSending message to - \n\t\tTopic: %v \n\t\tMessage: %s\n", topics, message)
	response, err := zbCli.NewMessage(topicMessges)
	if err != nil {
		log.Fatalf("NewMessage failed with error: %v", err)
	}
	return response
}

func sanitiseString(in string) string {
	return strings.Trim(in, "\n")
}

// shutDown terminates interactiveProducer session
func shutDown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()
}
