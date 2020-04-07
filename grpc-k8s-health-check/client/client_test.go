package client

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ZB-io/zbio/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
)

var (
	zbCli *Client
)

func TestMain(m *testing.M) {

	cfg := Config{
		// To run in Kubernetes change localhost => zbio-service
		Name:            "TestClient",
		ServiceEndPoint: "zbio-service:50002",
	}
	cli, err := New(cfg)

	if err != nil {
		log.Infof("Failed to get client, err=", err)
	} else {
		zbCli = cli
		os.Exit(m.Run())
	}
}

func TestHello(t *testing.T) {

	// Contact the server and print out its response.
	name := "Myname"

	resp, err := zbCli.ZbTest(name)
	if err != nil {
		t.Fatalf("could not greet: %v", err)
	}
	t.Logf("Response: %s", resp)
}

func TestCreateMultipleTopics(t *testing.T) {
	var topic string
	var topics []string
	for i := 10; i < 20; i++ {
		topic = fmt.Sprintf("test-topic-%d", i)
		t.Logf("Topic Name: %v", topic)
		topics = append(topics, topic)
	}
	ok, err := zbCli.CreateTopics(topics, "", 3, 1, 100000)
	if !ok {
		t.Errorf("Unable to create Topics due to error: %v", err)
	}
}

func TestCreateDuplicateTopics(t *testing.T) {
	var topic string
	var topics []string

	for i := 1; i < 10; i++ {
		topic = fmt.Sprintf("test-topic-%d", i)
		t.Logf("Topic Name: %v", topic)
		topics = append(topics, topic)
	}
	ok, err := zbCli.CreateTopics(topics, "", 3, 1, 100000)
	if !ok {
		t.Errorf("Unable to create duplicate Topics due to error: %v", err)
	}
}

func TestCreateInvalidTopics(t *testing.T) {
	var topic string
	var topics []string

	for i := 10; i < 12; i++ {
		topic = fmt.Sprintf("test@topic&%d", i)
		t.Logf("Topic Name: %v", topic)
		topics = append(topics, topic)
	}
	ok, err := zbCli.CreateTopics(topics, "", 3, 1, 100000)
	if !ok {
		t.Errorf("Unable to create Topics due to error: %v", err)
	}

	/*
		ok, err = zbCli.CreateTopics(topics, "", 2, 1, 1000)
		if !ok {
			t.Errorf("Unable to create duplicate Topics due to error: %v", err)
		}
	*/
}

func TestCreateSingleTopic(t *testing.T) {
	id := uuid.New()
	topic := fmt.Sprintf("topic-%s", id.String())
	ok, err := zbCli.CreateTopic(topic, "", 3, 0, 100000)
	if !ok {
		t.Errorf("Unable to create Topic %s due to error: %v", topic, err)
	}
}

func TestNewMessage(t *testing.T) {
	id := uuid.New()
	msg1 := fmt.Sprintf("Sample msg 1 for test-topic-1, i.e. hardcoded topic as of now. This message has a unique id generated for distinctness: %d", id)
	msg2 := fmt.Sprintf("Sample msg 2 for test-topic-1, i.e. hardcoded topic as of now. This message has a unique id generated for distinctness: %d", id)
	msg3 := fmt.Sprintf("Sample msg 3 for test-topic-1, i.e. hardcoded topic as of now. This message has a unique id generated for distinctness: %d", id)
	messages := []Message{
		Message{
			TopicName:     "test-topic-1",
			HintPartition: "",
			Data:          []byte(msg1),
		},
		Message{
			TopicName:     "test-topic-1",
			HintPartition: "",
			Data:          []byte(msg2),
		},
		Message{
			TopicName:     "test-topic-1",
			HintPartition: "",
			Data:          []byte(msg3),
		},
	}

	// log.Infof("%v\n", messages)
	response, err := zbCli.NewMessage(messages)
	if err != nil {
		t.Fatalf("NewMessage failed with error: %v", err)
	}
	log.Infof("NewMessage Response: %v", response)
}

func TestListTopic(t *testing.T) {
	list, err := zbCli.ListTopic(&empty.Empty{})

	if err != nil {
		t.Errorf("Error in Listing Topics %v", err)
	}

	for _, value := range list.GetTopics() {
		log.Infof("These are the Topics %v", value)
	}

}

func TestPeekMessage(t *testing.T) {

	topics := "test-topic-1"

	resp, err := zbCli.PeekMessages("testClient", "TestClientGroup", topics)

	if err != nil {
		t.Fatal("Peek message failed" + err.Error())
	} else {
		log.Infof("Response: %+v", resp)
	}
}

func TestReadMessage(t *testing.T) {

	topics := []string{"test-topic-1"} // , "test-topic-11"}

	msgChanMap, err := zbCli.ReadMessages("testClient", "TestClientGroup",
		topics)

	if err != nil {
		t.Fatal("Read message failed " + err.Error())
	}

	endCh := make(chan struct{}, 2)
	go func() {
		for {
			select {
			case msg := <-msgChanMap[topics[0]]:
				log.Infof("Message Received for topic %s Msg %s ", topics[0], string(msg))

				//	case msg := <-msgChanMap[topics[1]]:
				//		log.Infof("Message Received ", topics[1], string(msg))

			case <-endCh:
			}
		}
	}()

	time.Sleep(20 * time.Second)
	endCh <- struct{}{}

}

func TestDescribeTopic(t *testing.T) {
	list, err := zbCli.DescribeTopic([]string{"test-topic-10", "test-topic-11"})

	if err != nil {
		t.Errorf("Error in Describing Topic %v", err)
	}

	for _, value := range list.GetTopics() {
		log.Infof("These are the topic descriptions %v", value)
	}
}

//deleteTopic
func TestDeleteTopics(t *testing.T) {
	// var topic string
	var topic string
	var topics []string

	//creating array to delete the topics
	log.Infof("Creating an array of topics to delete: \n")
	for i := 12; i < 17; i++ {
		topic = fmt.Sprintf("test-topic-%d", i)
		topics = append(topics, topic)
	}
	log.Println(topics)

	ok, err := zbCli.DeleteTopic(topics) //array of topic is passed
	if !ok {
		t.Errorf("Unable to Delete Topics due to error: %v", err)
	}
	log.Infof("Topics Deleted.\n")
}

/*
func BenchmarkDeleteTopics(b *testing.B) {
	var t testing.T
	for n := 0; n < b.N; n++ {
		TestCreateTopics(&t)
		TestDeleteTopics(&t)
	}
}
*/
func BenchmarkCreateTopic(b *testing.B) {
	var t testing.T
	for n := 0; n < b.N; n++ {
		TestCreateSingleTopic(&t)
	}
}

// func TestListBroker(t *testing.T) {
// 	//log.Infof("")
// 	log.Infof("Broker List")
// 	//log.Infof("till here")
// 	list, err := zbCli.ListBroker()
// 	if err != nil {
// 		//log.Errorf("Error in Listing Brokers: %v", err)
// 	}
// 	for _, value := range list.GetBrokers() {
// 		log.Infof("Brokers --> %v", value)
// 	}

// }
