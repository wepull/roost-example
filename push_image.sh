#!/bin/bash
 ROOST_DIR=/var/tmp/Roost
 ZKE_NODE_DIR=$ROOST_DIR/nodes
 SSHOPTS="-i $ROOST_DIR/.ssh/id_rsa -o ConnectTimeout=5 -o LogLevel=ERROR -o StrictHostKeyChecking=false -o UserKnownHostsFile=/dev/null"
 IMAGE=$1
 IMAGE_TAR="$$.tar"

 if [ "$HOST" != "roost-utility" ]; then
     echo "Skip loading image to ZKE"
 else
 docker save $IMAGE > $IMAGE_TAR
 /usr/bin/find $ZKE_NODE_DIR -mindepth 1 -maxdepth 1 -type d -not -name tmp | while read node;
 do
    myprint "Copying $IMAGE_TAR  to $node"
    echo "scp $SSHOPTS $IMAGE_TAR $USER@$node:/tmp"
    scp $SSHOPTS $IMAGE_TAR $USER@$node:/tmp
    cmd_to_execute="docker load < /tmp/$IMAGE_TAR"
    ssh $SSHOPTS $USER@$node $cmd_to_execute
  done
 fi
