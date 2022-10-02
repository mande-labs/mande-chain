

# Kill docker containers
docker kill $(docker ps -q)

# Bring up the containers
(docker compose up) &

# Wait until the containers are up
sleep 10

# Add new validator -> test-node-2
docker cp test-node-1:/root/.mande-chain/config/genesis.json genesis.json
docker exec test-node-2 mkdir -p /root/.mande-chain/config
docker cp genesis.json test-node-2:/root/.mande-chain/config/genesis.json
rm -rf genesis.json

# Bring down the containers
# docker compose down
