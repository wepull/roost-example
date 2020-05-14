#!/bin/bash
 exit 0 # disable any image push from here
 USER_HOME=$(egrep "^USER_HOME" ~/.roost_user_env | awk -F'=' '{print $2}')
 SSHOPTS="-i $USER_HOME/.ssh/roost_id_rsa -o ConnectTimeout=5 -o LogLevel=ERROR -o StrictHostKeyChecking=false -o UserKnownHostsFile=/dev/null"
 IMAGE=$1
 IMAGE_TAR="$$.tar"

 if [ "$HOSTNAME" != "roost-utility" ]; then
     echo "Skip loading image to ZKE"
 else
 docker save $IMAGE > $IMAGE_TAR
 for node in `kubectl get nodes -o name | sed -e 's#node/##g' | grep -v master` # | while read node;
 do
    echo "scp $SSHOPTS $IMAGE_TAR $USER@$node:/tmp"
    scp $SSHOPTS $IMAGE_TAR $USER@$node:/tmp
    cmd_to_execute="docker load < /tmp/$IMAGE_TAR"
    echo "ssh $SSHOPTS $USER@$node $cmd_to_execute"
    ssh $SSHOPTS $USER@$node $cmd_to_execute
  done
 fi
 rm -f $IMAGE_TAR
