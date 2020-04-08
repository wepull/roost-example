/*
 * Copyright 2019 American Express Travel Related Services Company, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"log"
	"fmt"
	zb "github.com/ZB-io/zbio/client"
)

var (
	//log      *logrus.Logger
	zbclient *zb.Client
)

const (
	listenPort          = "5050"
	usdCurrency         = "USD"
	zbioServiceEndpoint = "zbio-service:50002"
	topicName           = "checkoutservice"
	zbioEnabled         = true
)

func getZBClient() (*zb.Client, error) {
	var err error
	fmt.Printf("Hello World12")
	if zbclient == nil && zbioEnabled {
		zbClientConfig := zb.Config{Name: "PlaceOrder", ServiceEndPoint: zbioServiceEndpoint}

		fmt.Printf("Hello World")
		zbclient, err = zb.New(zbClientConfig)
		fmt.Printf("Hello World1")
		if err != nil {
			fmt.Println("failed getting zbio client, errror: %+v", err)
			return nil, err
		}
	}
	return zbclient, nil
}

func initZBIO(str string) {
		zbclient, _ := getZBClient()
		if zbclient != nil {
			topicCreated, err := zbclient.CreateTopic(topicName, "", int32(1), int32(1), int32(1))
			fmt.Println(topicCreated)
			if err != nil {
				fmt.Println("failed to create topic, error: $v", err)
			}
			fmt.Println("create topic status: %s : %v", topicName, topicCreated)
		}

		var zbMessages []zb.Message
		zbMessages = append(zbMessages, zb.Message{
			TopicName:     topicName,
			Data:          []byte(fmt.Sprintf(str)),
			HintPartition: "",
		})
		sendMessageToZBIO(zbMessages)

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

// connectDB mimics a dummy database that waits some time and then changes the isDatabaseReady flag to true.
// This service is used to later check the readiness of the server.
func CollectClientLogs(str string)  {
	
	initZBIO(str)
	log.Println("Collected data from client %v",str)
}

func CollectServerLogs(str string)  {
	initZBIO(str)
	log.Println("Collected data from client %v",str)
}
