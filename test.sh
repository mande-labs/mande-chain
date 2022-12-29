

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


7d7e1953f5711048d09d1642a122ad9ff3d7e258@192.168.10.1:26656


5b6a072cd65650d168318343f61f6be58814e1a9@172.20.0.2:26656

mandeNode init test-node-2 --chain-id mande-testnet-1

mandeNode tx staking create-validator --amount="0cred" --pubkey=$(mandeNode tendermint show-validator) --from validator_2 --chain-id="mande-testnet-1"

wget 172.20.0.2:26657/genesis

docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' test-node-1

docker rm -f $(docker ps -a -q)
docker volume rm $(docker volume ls -q)


mandeNode tx voting create-vote cosmos1zxeknsu29kt7gn2ecfpyn9n6sma9tfnyqrqmrv 10000000 1 --from faucet1 --keyring-backend test --chain-id mande-testnet-1


# Bring down the containers
# docker compose down
