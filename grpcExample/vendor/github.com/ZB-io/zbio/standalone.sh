# Add host ip-addr to /etc/hosts as zbio-service and broker-0
# Example below:
#192.168.29.134  zbio-service
#192.168.29.134  broker-0
#set -x
start_zbserver() 
{
  echo "Starting ZBSERVER"
  ZB_PERSIST="false" SERVER_CRT="$GOPATH/src/github.com/ZB-io/zbio/security/cert/server.crt" SERVER_KEY="$GOPATH/src/github.com/ZB-io/zbio/security/cert/server.key" AUTO_DEPLOY_BROKERS="false" service/cmd/zbserver > /tmp/zbserver.log 2>&1 &
}

start_zbbroker() 
{
    if [ -z $1 ];then
      broker_name="broker-0"
    else
      echo "$1"
      broker_name=$1
    fi
    port=50003
    for n in 0 1 2 3 4 5
    do
      broker_name="broker-$n"
      x=`expr $n / 3`
      broker_port=`expr $port + $n`
      echo "Starting ZBBROKER : $broker_name for zbgroups-$x"
      BROKER_GROUP_NAME="zbgroups-$x" BROKER_PORT=$broker_port ZB_SVC_AP="localhost:50001" BROKER_NAME="${broker_name}" broker/cmd/zbbroker > /tmp/${broker_name}.log 2>&1 &
    done
}

test_single_topic()
{
  echo "Starting ZBCLIENT - CreateTopic"
  SERVER_CRT="$GOPATH/src/github.com/ZB-io/zbio/security/cert/server.crt" client/client.test -test.run TestCreateSingleTopic
  echo "================Done with CreateTopic ==========" 
}

test_topics()
{
  echo "Starting ZBCLIENT - CreateTopic"
  SERVER_CRT="$GOPATH/src/github.com/ZB-io/zbio/security/cert/server.crt" client/client.test -test.run CreateSingleTopic -test.count 1000
  echo "================Done with CreateTopic ==========" 
}

test_message()
{
  echo "Starting ZBCLIENT - Send Message"
  SERVER_CRT="$GOPATH/src/github.com/ZB-io/zbio/security/cert/server.crt" client/client.test -test.run TestNewMessage
  echo "Starting ZBCLIENT - Peek Message"
  SERVER_CRT="$GOPATH/src/github.com/ZB-io/zbio/security/cert/server.crt" client/client.test -test.run TestPeekMessage
  echo "================Done with SendMessage ==========" 
}

cleanup() {
  echo "Kill ZBSERVER and ZBBROKER"
  ps -aef | egrep 'zbserver|zbbroker' | egrep -v 'grep'
  `ps -aef | egrep 'zbserver|zbbroker' | egrep -v 'grep' | awk '{print $2}' | xargs kill -9`
}

main() 
{
    make build
    clear
    # for a in cleanup start_zbserver start_zbbroker test_single_topic test_topics test_message
    for a in cleanup start_zbserver start_zbbroker
    do
	$a
	sleep 2
    done
    for i in /tmp/*.log
    do
	echo "Log file ${i}"
        cat "${i}"
	echo "============================"
    done
    echo "Done"
}

if [ -z $@ ]; then
   main
else 
  echo $1
  $1
fi
