#!/bin/bash

if [ "$1" != "" ]; then
    echo "Config portal v3 contract with $1"
else
    exit
fi

if [ "$2" != "" ]; then
    echo "Config portal v3 with etherum node ip: $2"
else
    exit
fi

# export PORTAL_CONTRACT=$1
echo "PORTAL_CONTRACT=$1">>/etc/bash.bashrc
echo "GETH_NAME=$2">>/etc/bash.bashrc
source /etc/bash.bashrc

# run incognito chain
cd /go/incognito-chain && go build -o incognito

cd /go/incognito-chain && ./run_node.sh shard0-0 &
cd /go/incognito-chain && ./run_node.sh shard0-1 &
cd /go/incognito-chain && ./run_node.sh shard0-2 &
cd /go/incognito-chain && ./run_node.sh shard0-3 &

cd /go/incognito-chain && ./run_node.sh shard1-0 &
cd /go/incognito-chain && ./run_node.sh shard1-1 &
cd /go/incognito-chain && ./run_node.sh shard1-2 &
cd /go/incognito-chain && ./run_node.sh shard1-3 &

cd /go/incognito-chain && ./run_node.sh beacon-0 &
cd /go/incognito-chain && ./run_node.sh beacon-1 &
cd /go/incognito-chain && ./run_node.sh beacon-2 &
cd /go/incognito-chain && ./run_node.sh beacon-3 &
cd /go/incognito-highway && ./run.sh local