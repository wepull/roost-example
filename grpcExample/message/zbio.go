package message

import (
	"log"
	"os"

	zb "github.com/ZB-io/zbio/client"
)

const (
	// TopicName to be created in zbio
	TopicName string = "grpc-healthcheck-example"
	// ZBIOEnabled specifies whether zbio messaging is enabled
	zbioEnabled bool = true
)

var (
	zbserviceEndpoint string = "zbio-service:50002"
	zbclient          *zb.Client
)

// Config sets client configuration
func Config(name string) zb.Config {
	if zbsvc := os.Getenv("SERVICE_ADDRESS"); zbsvc != "" {
		zbserviceEndpoint = zbsvc
	}
	return zb.Config{Name: name, ServiceEndPoint: zbserviceEndpoint}
}

// GetZBClient returns client handler
func GetZBClient(cfg zb.Config) (*zb.Client, error) {
	var err error
	if zbclient == nil && zbioEnabled {
		zbclient, err = zb.New(cfg)
		if err != nil {
			log.Printf("failed getting zbio client, error: %+v", err)
			return nil, err
		}
	}
	return zbclient, nil
}

// InitZBIO initializes zbio client and creates a topic
func InitZBIO(cfg zb.Config) {
	zbclient, _ := GetZBClient(cfg)
	if zbclient != nil {
		topicCreated, err := zbclient.CreateTopic(TopicName, "", int32(1), int32(0), int32(10000))
		if err != nil {
			log.Printf("failed to create topic, error: %v", err)
		}
		log.Printf("create topic status: %s : %v", TopicName, topicCreated)
	}
}

// SendMessageToZBIO sends message to zbio
func SendMessageToZBIO(messages []zb.Message) {
	// send messages only if topic exists zbClient.DescribeTopics([]string{topicName})
	var topicFound = true
	if topicFound {
		newMessageStatus, err := zbclient.NewMessage(messages)
		if err != nil {
			log.Printf("failed to write message to zbio, error: %v", err)
		}
		for resp, status := range newMessageStatus {
			// msgInfo := strings.Split(resp, ",") // topicName:partition:messageIndex
			log.Printf("messages sent to zbio status? TopicName: %s\t Status: %s", resp, status)
		}
	}
}
