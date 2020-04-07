#"go","test","-v","-count=1","-timeout=30s","/go/src/github.com/ZB-io/zbio/client","-run=TestCreateTopic"
#go test -v /go/src/github.com/ZB-io/zbio/client
#go test -v -count=1 -timeout=30s /go/src/github.com/ZB-io/zbio/client -run=TestCreateTopic
echo "Calling ZBClient"
#export SERVER_CRT=/zbclient/security/cert/server.crt && ./client.test -test.run TestCreateTopic
export SERVER_CRT=/zbclient/security/cert/server.crt && ./client.test -test.run CreateSingleTopic -test.count 1000
