#!/bin/sh

CHAIN_ID=mande-testnet-1
MONIKER=test-node-2
VALIDATOR_NAME=validator_2

## Initialize the node
echo "Initializing the node $MONIKER, chain id: $CHAIN_ID"
mandeNode init $MONIKER --chain-id $CHAIN_ID

## Add keys
mandeNode keys add $VALIDATOR_NAME --keyring-backend test

## Modify the config
CONFIG_TOML=/root/.mande-chain/config/config.toml
APP_TOML=/root/.mande-chain/config/app.toml
if [[ -z "$(grep '\[api\]' -A 5 $APP_TOML | grep true)" ]]; then
  sed -i '1,/enable = false/{s/enable = false/enable = true/g}' $APP_TOML
fi
sed -i 's/cors_allowed_origins.*/cors_allowed_origins = ["*"]/g' $CONFIG_TOML
sed -i 's/cors_allowed_methods.*/cors_allowed_methods = ["*"]/g' $CONFIG_TOML
sed -i 's/cors_allowed_headers.*/cors_allowed_headers = ["*"]/g' $CONFIG_TOML

## Get node 1 peer id
ID_NODE_1=$(docker exec test-node-1 mandeNode tendermint show-node-id)
IP_NODE_1=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' test-node-1)
PORT_NODE_1="26656"
PEER_ID_NODE_1="$ID_NODE_1@$IP_NODE_1:$PORT_NODE_1"
echo "Node1 Peer ID: $PEER_ID_NODE_1"

## Set node 1 peer id in node 2 persistent peers
sed -i '/^persistent_peers =/s/=.*/=$PEER_ID_NODE_1/' $CONFIG_TOML
cat $CONFIG_TOML

## Start the node in the background
echo "Starting the node..."
mandeNode start --pruning=nothing --rpc.laddr=tcp://0.0.0.0:26657 &


