

# Kill docker containers
#docker kill $(docker ps -q)

# Bring up the containers
#(docker compose up) &

# Wait until the containers are up
#sleep 10

# Add new validator -> test-node-2
MONIKER=test-node-2
CHAIN_ID=mande-testnet-1
docker exec test-node-2 mandeNode init ${MONIKER} --chain-id ${CHAIN_ID}

docker cp test-node-1:/root/.mande-chain/config/genesis.json genesis.json
docker exec test-node-2 mkdir -p /root/.mande-chain/config
docker cp genesis.json test-node-2:/root/.mande-chain/config/genesis.json
rm -rf genesis.json

ID_NODE_1=$(docker exec test-node-1 mandeNode tendermint show-node-id)
echo "Node1 ID: $ID_NODE_1"

IP_NODE_1="192.168.10.1"
PORT_NODE_1="26656"

PEER_ID_NODE_1="$ID_NODE_1@$IP_NODE_1:$PORT_NODE_1"
echo "Node1 Peer ID: $PEER_ID_NODE_1"

docker exec test-node-1 cat /root/.mande-chain/config/config.toml
docker exec test-node-1 sed -c -i "s/\(persistent_peers *= *\).*/\1$PEER_ID_NODE_1/" /root/.mande-chain/config/config.toml
docker exec test-node-1 cat /root/.mande-chain/config/config.toml




# Bring down the containers
# docker compose down
