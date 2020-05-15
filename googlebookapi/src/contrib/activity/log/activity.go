package log

import (
	"fmt"
	"log"
	"os"
	"strings"

	zb "github.com/ZB-io/zbio/client"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

var (
	//log      *logrus.Logger
	zbclient *zb.Client
)

const (
	listenPort  string = "5050"
	usdCurrency string = "USD"
	topicName   string = "googleBookAPI"
	zbioEnabled bool   = true
)

// Tip: USE ENV[SERVICE_ADDRESS] to set service endpoint
var zbioServiceEndpoint string = "zbio-service:50002"

func init() {
	_ = activity.Register(&Activity{}, New)
	initZBIO()
}

type Input struct {
	Message    string `md:"message"`    // The message to log
	AddDetails bool   `md:"addDetails"` // Append contextual execution information to the log message
	UsePrint   bool   `md:"usePrint"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message":    i.Message,
		"addDetails": i.AddDetails,
		"usePrint":   i.UsePrint,
	}
}

func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance
	return act, nil
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.Message, err = coerce.ToString(values["message"])
	if err != nil {
		return err
	}
	i.AddDetails, err = coerce.ToBool(values["addDetails"])
	if err != nil {
		return err
	}

	i.UsePrint, err = coerce.ToBool(values["usePrint"])
	if err != nil {
		return err
	}

	return nil
}

var activityMd = activity.ToMetadata(&Input{})

// Activity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func getZBClient() (*zb.Client, error) {
	var err error
	if zbclient == nil && zbioEnabled {
		if zbsvc := os.Getenv("SERVICE_ADDRESS"); zbsvc != "" {
			zbioServiceEndpoint = zbsvc
		}
		log.Printf("service endpoint is: %s\n", zbioServiceEndpoint)
		zbClientConfig := zb.Config{Name: "GoogleBookAPI", ServiceEndPoint: zbioServiceEndpoint}

		zbclient, err = zb.New(zbClientConfig)
		if err != nil {
			log.Printf("failed getting zbio client, errror: %+v\n", err)
			return nil, err
		}
	}
	return zbclient, nil
}

func initZBIO() {
	zbclient, _ := getZBClient()
	if zbclient != nil {
		topicCreated, err := zbclient.CreateTopic(topicName, "", int32(1), int32(1), int32(10000))
		if err != nil {
			log.Printf("failed to create topic, error: %v\n", err)
		}
		log.Printf("create topic status: TopicName: %s\tStatus: %v\n", topicName, topicCreated)
	}
}

func sendMessageToZBIO(ctx activity.Context, messages []zb.Message) {
	// send messages only if topic exists zbClient.DescribeTopics([]string{topicName})
	var topicFound = true
	if topicFound {
		newMessageStatus, err := zbclient.NewMessage(messages)
		if err != nil {
			ctx.Logger().Errorf("failed to write message to zbio, error: %v", err)
		}
		for resp, status := range newMessageStatus {
			msgInfo := strings.Split(resp, ",")
			ctx.Logger().Infof("messages sent to zbio status? TopicName: %s\t Status: %s", msgInfo[0], status)
		}
	}
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)

	msg := input.Message

	if input.AddDetails {
		msg = fmt.Sprintf("'%s' - HostID [%s], HostName [%s], Activity [%s]", msg,
			ctx.ActivityHost().ID(), ctx.ActivityHost().Name(), ctx.Name())
	}

	zbMessages := []zb.Message{
		zb.Message{
			TopicName:     topicName,
			Data:          []byte(fmt.Sprintf(msg)),
			HintPartition: "",
		},
	}
	sendMessageToZBIO(ctx, zbMessages)

	if input.UsePrint {
		log.Println("UsePrint message: ", msg)
	} else {
		ctx.Logger().Info(msg)
	}
	return true, nil
}
