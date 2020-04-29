#set -x 
cleanup()
{
  kubectl delete -f kubernetes-manifests/zbio.yaml 
  kubectl delete -f kubernetes-manifests/zbio-client.yaml 
}

deploy_zbsvc_zbbroker()
{
  kubectl create -f kubernetes-manifests/zbio.yaml 
}

deploy_zbclient()
{
  kubectl delete -f kubernetes-manifests/zbio-client.yaml 
  kubectl create -f kubernetes-manifests/zbio-client.yaml 
  sleep 5
  kubectl logs -f zbio-client
}

show_pods()
{
    kubectl get pods | egrep 'NAME|Running|Completed'
}

show_logs()
{
    for pod in `kubectl get pods | egrep 'Running|Completed' | awk '{print $1}' | grep -v 'NAME|zbio-service'`
    do
	echo "------------------------------------------------------"
	echo "LOG for ${pod}"
	kubectl logs ${pod}
	echo
	echo "------------------------------------------------------"
	echo
    done
    service_pod=`kubectl get pods | grep 'zbio-service' | awk '{print $1}'`
    kubectl logs -f ${service_pod}
}

main() 
{
    clear
    make build-linux && make image
    # for a in cleanup deploy_zbsvc_zbbroker deploy_zbclient show_pods show_logs
    for a in cleanup deploy_zbsvc_zbbroker show_pods show_logs
    do
	$a
	sleep 3
    done
    echo "Done"
}

if [ -z $@ ]; then
   main
else 
  echo $1
  $1
fi
#main
