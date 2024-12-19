# ranger-hadoop-blockchain

## Blockchain configuration (Hyperledger fabric)

- To create crypto-config files file use:

```sh
./commands/cryptogen generate --config=./crypto-config.yaml
```

- To create genesis.block file use:

```sh
./commands/configtxgen -configPath=. -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID mychannel
```

- To create channel.tx file use:

```sh
./commands/configtxgen -configPath=. -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel
```

finally, finally, to build up the container run:

```sh
docker-compose -f docker-compose-cli.yaml up -d
```
