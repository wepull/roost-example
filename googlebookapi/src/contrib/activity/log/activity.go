package log

import (
	"fmt"
	//"log"

	zb "github.com/ZB-io/zbio/client"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

var (
	//log      *logrus.Logger
	zbclient *zb.Client
)

const (
	listenPort          = "5050"
	usdCurrency         = "USD"
	zbioServiceEndpoint = "zbio-service:50002"
	topicName           = "googleBookAPI"
	zbioEnabled         = true
)

func init() {
	_ = activity.Register(&Activity{}, New)
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

	/*s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)*/
	/*if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}*/

	//cmd := exec.Command("man ls")

	act := &Activity{} //add aSetting to instance
	fmt.Println("Hello World")
	initZBIO("")
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
		zbClientConfig := zb.Config{Name: "PlaceOrder", ServiceEndPoint: zbioServiceEndpoint}

		zbclient, err = zb.New(zbClientConfig)
		if err != nil {
			fmt.Println("failed getting zbio client, errror: %+v", err)
			return nil, err
		}
	}
	return zbclient, nil
}

func initZBIO(str string) {
	if str == "" {
		zbclient, _ := getZBClient()
		if zbclient != nil {
			topicCreated, err := zbclient.CreateTopic(topicName, "", int32(1), int32(1), int32(10000))
			fmt.Println(topicCreated)
			if err != nil {
				fmt.Println("failed to create topic, error: $v", err)
			}
			fmt.Println("create topic status: %s : %v", topicName, topicCreated)
		}
	} else {
		var zbMessages []zb.Message
		zbMessages = append(zbMessages, zb.Message{
			TopicName:     topicName,
			Data:          []byte(fmt.Sprintf(str)),
			HintPartition: "",
		})
		sendMessageToZBIO(zbMessages)
	}

}

func sendMessageToZBIO(messages []zb.Message) {
	// send messages only if topic exists zbClient.DescribeTopics([]string{topicName})
	var topicFound = true
	if topicFound {
		fmt.Println(messages)
		newMessageStatus, err := zbclient.NewMessage(messages)
		if err != nil {
			fmt.Println("failed to write message to zbio, error:", err)
		}
		fmt.Println("messages sent to zbio, %v", newMessageStatus)
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

	fmt.Println("Hello Message")
	initZBIO(msg)

	if input.UsePrint {
		fmt.Println(msg)
	} else {
		ctx.Logger().Info(msg)
	}
	return true, nil
}
