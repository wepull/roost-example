//go:generate protoc -I ../rpc/common --go_out=plugins=grpc:$GOPATH/src ../rpc/common/cloud_event.proto
//go:generate protoc -I ../rpc/common --go_out=plugins=grpc:$GOPATH/src ../rpc/common/common.proto
//go:generate protoc -I ../rpc --go_out=plugins=grpc:$GOPATH/src ../rpc/service.proto

package client

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"path"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/ZB-io/zbio/client/config"
	"github.com/ZB-io/zbio/log"
	pbcommon "github.com/ZB-io/zbio/rpc/common"
	pbservice "github.com/ZB-io/zbio/rpc/service"
	zbgrpctracer "github.com/ZB-io/zbio/service/middleware/tracer"
	zbtracer "github.com/ZB-io/zbio/tracer"
	"github.com/ZB-io/zbio/util"

	//TODO: Instead of using service/middle/tracing either set middleware for every microservices or at zbio level

	"github.com/golang/protobuf/ptypes/empty"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	otkey "go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	rootDir = clientRootDir()
	tracer  = opentracing.GlobalTracer()
	// MaxMsgSize is maximum allowed size for a test message
	MaxMsgSize      = 64 * 1024 * 1024 // 64 MB default
	msgReadInterval = 200 * time.Millisecond
)

type Config struct {
	Name            string
	ServiceEndPoint string

	// Auth related
	Scheme string
	Token  string
}

type Client struct {
	Config
	conn  *grpc.ClientConn
	zbCli pbservice.ServiceClient
}

type Broker_BrokerState int32

type BrokerInfo struct {
	Name      string             ` json:"name,omitempty"`
	Groupname string             ` json:"groupname,omitempty"`
	State     Broker_BrokerState `json:"state,omitempty"`
}

type Consumer struct {
	clientId string
	groupId  string
	topics   []string

	errorCh     chan error
	rebalanceCh chan *pbcommon.RebalanceResponse

	partitionMap     map[string]*Partition
	partitionMapLock sync.RWMutex

	msgChanMap map[string]chan []byte
	msgChaLock sync.RWMutex

	client    *Client
	sessionId string
}

type Partition struct {
	key      string
	topic    string
	partNum  int32
	offset   int64
	msgChan  chan []byte
	stopCh   chan struct{}
	consumer *Consumer
}

/*
const (
	Broker_RUNNING = pbcommon.Broker_READY
	Broker_DOWN    = pbcommon.Broker_WAITING
	Broker_FAILED  = pbcommon.Broker_UNHEALTHY
)

type BrokerGroupInfo struct {
	Name    string
	Brokers []*BrokerInfo ` json:"brokers,omitempty"`
}
*/
type Message struct {
	TopicName     string `json:"topic"`
	Data          []byte `json:"messages"`
	HintPartition string `json:"hint"`
}

// MessageResponse is returned to client with topic, messages_sent to those topics
type MessageResponse struct {
	// Messages stored are {topic1:5,topic2:3}
	status []map[string]int32
}

func New(cfg Config) (*Client, error) {
	config.InitTrace()

	log.Debugf("clientName is: %s\tzbio service endpoint is: %s\n", cfg.Name, cfg.ServiceEndPoint)

	//example with a 100ms linear backoff interval, and retry only on NotFound and Unavailable.
	opts := []grpcretry.CallOption{
		grpcretry.WithBackoff(grpcretry.BackoffExponential(100 * time.Millisecond)),
	}

	// Create the client TLS credentials
	TLSConfig := &tls.Config{
		ServerName:             "zb.io",
		Certificates:           []tls.Certificate{tls.Certificate{}},
		Rand:                   rand.Reader,
		SessionTicketsDisabled: false,
		MinVersion:             tls.VersionTLS12,
		InsecureSkipVerify:     true,
	}
	creds := credentials.NewTLS(TLSConfig)

	// Set up a connection to the server.
	conn, err := grpc.DialContext(GetClientContext(cfg),
		cfg.ServiceEndPoint,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(MaxMsgSize),
			grpc.MaxCallSendMsgSize(MaxMsgSize),
		),
		grpc.WithUnaryInterceptor(grpcmiddleware.ChainUnaryClient(
			// otgrpc.OpenTracingClientInterceptor(tracer),
			grpcretry.UnaryClientInterceptor(opts...),
			zbgrpctracer.UnaryClientInterceptor,
		)),
		grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer)))

	if err != nil {
		log.Fatalf("couldn't connect %v", err)
		return nil, err
	}
	return &Client{
		conn:   conn,
		Config: cfg,
		zbCli:  pbservice.NewServiceClient(conn),
	}, nil
}

func clientRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func (cli *Client) Close() {
	defer cli.conn.Close()
}

func (cli *Client) ZbTest(str string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")
	//	ctx = opentracing.ContextWithSpan(ctx, span)
	defer cancel()

	/*
		span := tracer.StartSpan("requestAtClientSide", opentracing.StartTime(time.Now()))
		defer span.Finish()
	*/

	var span zbtracer.Span
	tr := zbtracer.New(zbtracer.TraceCfg{Name: "zbio.tracer.client.ZbTest"})
	ctx, span = tr.StartNewSpan(ctx, "client.ZbTest")
	defer span.End()

	r, err := cli.zbCli.TestMessage(ctx, &pbcommon.TestRequest{Name: str})
	stat, ok := status.FromError(err)
	if ok {
		return "", status.Errorf(stat.Code(), "Message: %s\t, Details: %v,\tError: %v", stat.Message(), stat.Details(), stat.Err())
	}
	log.Debugf("ZbTest newMessage Res: %v, err: %v, stat: %v, ok: %v", r, err, stat, ok)
	return r.GetMessage(), err
}

/*
func (cli *Client) GetBrokers() ([]*BrokerInfo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := cli.zbCli.ListBroker(ctx, &pb.Empty{})

	if err == nil && resp != nil {

		brokers := []*BrokerInfo{}
		for _, br := range resp.Brokers {
			brokers = append(brokers, &BrokerInfo{
				Name:      br.Name,
				Groupname: br.Group,
				State:     Broker_BrokerState(br.State),
			})
		}
		return brokers, err
	}

	return nil, err
}

func (cli *Client) ListBrokerGroup() ([]*pb.BrokerGroup, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := cli.zbCli.ListBrokerGroup(ctx, &pb.Empty{})

	if err == nil && resp != nil {
		return resp.Groups, err
	}

	return nil, err
}

func (cli *Client) GetBrokerGroup(name string) (*BrokerGroupInfo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.GetBrokerGroupRequest{
		Name: name,
	}
	resp, err := cli.zbCli.GetBrokerGroup(ctx, &req)

	if err != nil {
		return nil, err
	} else if resp == nil {
		return nil, errors.New("Got nil response")
	} else {

		brokers := []*BrokerInfo{}

		for brName, resp := range resp.Responses {
			brokers = append(brokers, &BrokerInfo{
				Name:      brName,
				Groupname: name,
				State:     Broker_BrokerState(resp.Code),
			})
		}

		return &BrokerGroupInfo{
			name,
			brokers,
		}, nil

	}
}
*/

// CreateTopic creates topic with "p" partitions
func (cli *Client) CreateTopic(topic string, brokerGroup string, p, r int32, retentionDuration int32) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	// create tracer
	methodName := util.CallerFunc(2)
	tr := zbtracer.NewTracer(zbtracer.TraceOptions{Name: methodName, Kind: zbtracer.TraceKindClient})
	ctx, span := tr.Start(ctx, methodName, trace.WithNewRoot())
	defer span.End()

	ok, err := cli.CreateTopics([]string{topic}, brokerGroup, p, r, retentionDuration)
	span.AddEventWithTimestamp(ctx, time.Now(), "status", []core.KeyValue{otkey.New("status").Bool(ok)}...)
	if err != nil {
		span.AddEventWithTimestamp(ctx, time.Now(), "status", []core.KeyValue{otkey.New("Error").String(err.Error())}...)
		log.Errorf("error for CreateTopic: %s", err.Error())
	}
	return true, nil
}

// CreateTopics creates topic with "p" partitions
func (cli *Client) CreateTopics(topicNames []string, brokerGroup string, p, r int32, retentionDuration int32) (bool, error) {

	log.Debugf("Creating Topic %v with %v Partitions", topicNames, p)
	topics := []*pbcommon.Topic{}
	// Create Trace

	for _, topic := range topicNames {
		// Checking of regular Expression
		// var validID = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
		ok := true // validID.MatchString(topic)

		if ok == true {
			topics = append(topics, &pbcommon.Topic{
				Name: topic,
				// BrokerGroup: brokerGroup,
				NumPartitions:               p,
				NumReplicas:                 r,
				MsgRetentionDurationSeconds: retentionDuration,
			})
		} else {
			log.Errorf("Invalid Topic Name, %v", topic)
		}
	}

	request := pbcommon.TopicRequest{
		Topics: topics,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")

	// create tracer
	methodName := util.CallerFunc(2)
	tr := zbtracer.NewTracer(zbtracer.TraceOptions{Name: methodName, Kind: zbtracer.TraceKindClient})
	ctx, span := tr.Start(ctx, methodName, trace.WithNewRoot())
	defer span.End()

	/*
		var span zbtracer.Span
		tr := zbtracer.New(zbtracer.TraceCfg{Name: createTracerName("createTopics")})
		ctx, span = tr.StartNewSpan(ctx, "client.createTopics")
		defer span.End()
		span.AddEventWithTimestamp(ctx, "CreateTopicClientRequest", []core.KeyValue{otkey.New("TopicCount").Int(len(topics))})
	*/

	span.AddEventWithTimestamp(ctx, time.Now(), "CreateTopicClientRequest", []core.KeyValue{otkey.New("TopicCount").Int(len(topics))}...)
	// Create ctx, span
	response, err := cli.zbCli.CreateTopic(ctx, &request)
	if err != nil {
		return false, err
	}
	for topic, resp := range response.Responses {
		span.AddEventWithTimestamp(ctx, time.Now(), "CreateTopicClientResponse",
			[]core.KeyValue{
				otkey.New("topicName").String(topic),
				otkey.New("response").String(resp.GetMessage())}...)

		//log.Debugf("Topic: %s\nBroker: %s\nPartition: %v\n", response.Response)
		log.Debugf("Topic Response: %v  %v", topic, resp)
	}
	return true, nil
}

// NewMessage writes messages to the topic and returns "n" numbers of messages which are written.
// 	Response format is: response[key]value
// 		key: is in the form (topicName:partitionHint:messageId) // from broker/broker.go
//		value: status.code().String()
//  TODO: (update key format if not meaningful)
func (cli *Client) NewMessage(messages []Message) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")

	// create tracer
	methodName := util.CallerFunc(2)
	tr := zbtracer.NewTracer(zbtracer.TraceOptions{Name: methodName, Kind: zbtracer.TraceKindClient})
	ctx, span := tr.Start(ctx, methodName, trace.WithNewRoot(), trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	newMessageResponse := make(map[string]string, len(messages))
	var pbMessages []*pbcommon.Message

	for i, msg := range messages {
		pbMessage := pbcommon.Message{
			Topic:         msg.TopicName,
			Data:          msg.Data,
			HintPartition: msg.HintPartition,
		}
		pbMessages = append(pbMessages, &pbMessage)
		log.Debugf("Message {%v}: %s, %s, %s", i, msg.TopicName, msg.HintPartition, msg.Data)
	}

	span.AddEventWithTimestamp(ctx, time.Now(), "NewMessage Request", []core.KeyValue{otkey.New("message count").Int(len(messages))}...)

	log.Debugf("Requested NewMessage data: %v", pbMessages)
	request := pbcommon.NewMessageRequest{
		Messages: pbMessages,
	}

	response, err := cli.zbCli.NewMessage(ctx, &request)
	if err != nil {
		return nil, err
	}
	responses := response.GetResponses()
	for t, r := range responses {
		newMessageResponse[t] = r.GetCode().String()
	}
	return newMessageResponse, nil
}

func ctxWithToken(ctx context.Context, scheme string, token string) context.Context {
	md := metadata.Pairs("authorization", fmt.Sprintf("%s %v", scheme, token))
	nCtx := metautils.NiceMD(md).ToOutgoing(ctx)
	return nCtx
}

func (cli *Client) ReadMessages(clientId, clientGroup string, topics []string) (map[string]chan []byte, <-chan error, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")

	// create tracer
	methodName := util.CallerFunc(2)
	tr := zbtracer.NewTracer(zbtracer.TraceOptions{Name: methodName, Kind: zbtracer.TraceKindClient})
	ctx, span := tr.Start(ctx, methodName, trace.WithNewRoot())
	defer span.End()

	span.AddEventWithTimestamp(ctx, time.Now(), "ReadMessages Request", []core.KeyValue{otkey.String("clientID", clientId)}...)
	stream, err := cli.zbCli.SubscribeConsumer(ctx)
	if err != nil {
		log.Errorf("Unable to subscribe consumer. Error: %v", err)
		return nil, nil, err
	}

	// Send subscribe request
	req := pbcommon.ConsumerRequest{
		Consumer: &pbcommon.Consumer{
			Name:          clientId,
			ConsumerGroup: clientGroup,
			Topics:        topics,
		},
	}
	stream.Send(&req)

	sessionId := ""
	if subscribeResp, err := stream.Recv(); err != nil {
		log.Errorf("Unable to receive over stream. Error: %v", err)
		return nil, nil, err
	} else {
		resp := subscribeResp.GetInitResponse().Responses
		for _, respStatus := range resp {
			if respStatus.Code != pbcommon.ResponseStatus_OK {
				return nil, nil, errors.New("Failed to create subscription " + respStatus.Message)
			}
		}
		sessionId = subscribeResp.GetInitResponse().SessionId
		span.AddEventWithTimestamp(ctx, time.Now(), "ReadMessages Consumer Response", []core.KeyValue{
			otkey.String("sessionID", sessionId),
			otkey.String("clientID", clientId),
		}...)

		log.Debugf("Received consumer init response %+v %+v", sessionId, resp)
	}

	msgChanMap := make(map[string]chan []byte)
	for _, topic := range topics {
		msgChanMap[topic] = make(chan []byte)
	}
	consumer := Consumer{
		clientId:         clientId,
		sessionId:        sessionId,
		groupId:          clientGroup,
		topics:           topics,
		client:           cli,
		errorCh:          make(chan error),
		rebalanceCh:      make(chan *pbcommon.RebalanceResponse),
		msgChanMap:       msgChanMap,
		msgChaLock:       sync.RWMutex{},
		partitionMap:     make(map[string]*Partition),
		partitionMapLock: sync.RWMutex{},
	}

	go func() {
		log.Debugf("%s Started reading rebalance messages ", consumer.clientId)
		defer log.Debugf("%s Ended reading rebalance messages ", consumer.clientId)
		for {
			in, err := stream.Recv()

			if err == io.EOF {
				consumer.errorCh <- err
				return
			}
			if err != nil {
				consumer.errorCh <- err
				return
			}

			// Got Re-balance message
			consumer.rebalanceCh <- in.GetRebalResponse()
			log.Debugf("Got %v\n", in.GetResponse())
		}
	}()
	go consumer.handleRebalance()
	return msgChanMap, consumer.errorCh, nil
}

func (cli *Client) PeekMessages(clientId, clientGroup, topic string) (*pbcommon.GetMessageResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")
	defer cancel()

	// Send subscribe request
	req := pbcommon.PeekMessageRequest{
		Info: &pbcommon.MessageInfo{
			TopicName: topic,
			Partition: 1,
			Offset:    1,
		},
	}

	return cli.zbCli.PeekMessage(ctx, &req)
}

func (cons *Consumer) handleRebalance() {

	for {
		rebalanceInfo := <-cons.rebalanceCh

		log.Debugf("Received rebalance for topic %s. New %+v Added %+v Deleted %+v",
			rebalanceInfo.TopicName, rebalanceInfo.Owned, rebalanceInfo.Added, rebalanceInfo.Deleted)

		/*
			for _, part := range rebalanceInfo.Added {
				log.Infof("Added Partition %+v offset %+v ", part.Partnum,  part.Offset)
			}

			for _, part := range rebalanceInfo.Deleted {
				log.Infof("Delete Partition %+v offset %+v ", part.Partnum,  part.Offset)

			}
		*/

		cons.partitionMapLock.RLock()
		for _, part := range rebalanceInfo.Deleted {
			log.Debugf("Giving up ownership of topic %s partition %d",
				rebalanceInfo.TopicName, part.Partnum)
			key := getPartitionKey(rebalanceInfo.TopicName, part.Partnum)

			if part, found := cons.partitionMap[key]; !found {
				log.Debugf("Partition is not found %s", key)
			} else {
				part.stopCh <- struct{}{}
			}
		}
		cons.partitionMapLock.RUnlock()

		msgCh := cons.msgChanMap[rebalanceInfo.TopicName]

		for _, part := range rebalanceInfo.Added {
			log.Debugf("Claiming ownership of topic %s partition %d",
				rebalanceInfo.TopicName, part.Partnum)
			partition := NewPartition(rebalanceInfo.TopicName, part.Partnum, part.Offset, msgCh, cons)
			go partition.Start(cons)
		}
	}
}

// NewPartition is going to
func NewPartition(topic string, partNum int32, offset int64, msgChan chan []byte, consumer *Consumer) *Partition {
	return &Partition{
		key:      getPartitionKey(topic, partNum),
		topic:    topic,
		partNum:  partNum,
		offset:   offset,
		msgChan:  msgChan,
		consumer: consumer,
		stopCh:   make(chan struct{}),
	}
}

func getPartitionKey(name string, num int32) string {

	return fmt.Sprintf("%s:%d", name, num)
}

func (part *Partition) Start(cons *Consumer) {

	log.Debugf("Started reading partition %s -> %d", part.topic, part.partNum)
	defer log.Debugf("Ended reading partition %s -> %d", part.topic, part.partNum)

	part.consumer.partitionMapLock.Lock()
	part.consumer.partitionMap[part.key] = part
	part.consumer.partitionMapLock.Unlock()

	readUt := util.Timer(msgReadInterval)
	defer readUt.Stop()

	for {

		select {
		case _ = <-part.stopCh:
			part.Close()
			return
		case <-readUt.Chan():
			// Make the call to get the message.
			// After getting the message pass it to part.msgChan
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			ctx = ctxWithToken(ctx, "bearer", "zb-test-token")
			defer cancel()

			// Send subscribe request
			req := pbcommon.MessageRequest{
				Info: &pbcommon.MessageInfo{
					TopicName: part.topic,
					Partition: part.partNum,
					Offset:    part.offset,
				},
				SessionId: cons.sessionId,
			}

			if resp, err := cons.client.zbCli.GetMessage(ctx, &req); err == nil {

				if resp != nil && resp.Response != nil && resp.Response.Code == pbcommon.ResponseStatus_OK {
					for _, msg := range resp.Messages {
						part.msgChan <- msg.Data
						part.offset = msg.Id + 1
					}
				}
			} else {
				log.Errorf("GetMessage Error : %+v ", err)
				cons.errorCh <- fmt.Errorf("Error reading message from partition. PartitionKey: %v, Error: %v", part.key, err)
			}
		}
	}
}

func (part *Partition) Close() {

	part.consumer.partitionMapLock.Lock()
	delete(part.consumer.partitionMap, part.key)
	part.consumer.partitionMapLock.Unlock()
}

// List Topics present in the topicHeapStore
func (cli *Client) ListTopic(*empty.Empty) (*pbcommon.DescribeTopicResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")

	// create tracer
	methodName := util.CallerFunc(2)
	tr := zbtracer.NewTracer(zbtracer.TraceOptions{Name: methodName, Kind: zbtracer.TraceKindClient})
	ctx, span := tr.Start(ctx, methodName, trace.WithNewRoot())
	defer span.End()

	span.AddEventWithTimestamp(ctx, time.Now(), "ListTopic Request")

	response, err := cli.zbCli.ListTopic(ctx, &empty.Empty{})
	if err != nil {
		span.AddEventWithTimestamp(ctx, time.Now(), "ListTopic Response", []core.KeyValue{otkey.String("Error", err.Error())}...)
		return nil, err
	}
	span.AddEventWithTimestamp(ctx, time.Now(), "ListTopic Response", []core.KeyValue{otkey.Int("topic count", len(response.GetTopics()))}...)
	return response, err
}

// DescribeTopic from topicHeapStore
func (cli *Client) DescribeTopic(topicNames []string) (*pbcommon.DescribeTopicResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")

	// create tracer
	methodName := util.CallerFunc(2)
	tr := zbtracer.NewTracer(zbtracer.TraceOptions{Name: methodName, Kind: zbtracer.TraceKindClient})
	ctx, span := tr.Start(ctx, methodName, trace.WithNewRoot())
	defer span.End()

	span.AddEventWithTimestamp(ctx, time.Now(), "DescribeTopic Request", []core.KeyValue{otkey.Int("topic count", len(topicNames))}...)

	request := pbcommon.DescribeTopicRequest{Names: topicNames}

	response, err := cli.zbCli.DescribeTopic(ctx, &request)
	if err != nil {
		return nil, err
	}

	return response, err
}

// DeleteTopic via topic name
func (cli *Client) DeleteTopic(topicName []string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")

	// create tracer
	methodName := util.CallerFunc(2)
	tr := zbtracer.NewTracer(zbtracer.TraceOptions{Name: methodName, Kind: zbtracer.TraceKindClient})
	ctx, span := tr.Start(ctx, methodName, trace.WithNewRoot())
	defer span.End()

	span.AddEventWithTimestamp(ctx, time.Now(), "DeleteTopic Request", []core.KeyValue{otkey.Int("topic count", len(topicName))}...)

	request := pbcommon.DeleteTopicRequest{
		Names: topicName,
	}

	response, err := cli.zbCli.DeleteTopic(ctx, &request)

	if err == nil {
		// log.Println(response)
		for _, del := range response.DeletedTopics {
			log.Debugf("Delete topic response: %v Code:%v Message:%v", del,
				response.Responses[del.Name].Code, response.Responses[del.Name].Message)
		}
		return true, nil
	}
	return false, err
}

// func (cli *Client) ListBroker() (*pb.ListBrokersResponse, error) {

// 	log.Infof("Itna to chala")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	//	ctx = ctxWithToken(ctx, "bearer", "zb-test-token")

// 	defer cancel()

// 	list, err := cli.zbCli.ListBroker(ctx, &empty.Empty{})
// 	log.Infof("yaha aya")

// 	if err != nil {
// 		log.Infof("Error tha")

// 		return nil, err
// 	}
// 	return list, err
// }

// GetClientContext gives context with users information
func GetClientContext(cfg Config) context.Context {
	return context.WithValue(context.Background(), "user_id", cfg.Name)
}

// CreateTracer creates client prefixed tracer
func CreateTracer(name string) trace.Tracer {
	return global.TraceProvider().Tracer("zbio.tracer.client." + name)
}

func createTracerName(name string) string {
	return "zbio.tracer.client." + name
}
