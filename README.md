# ranger-hadoop-blockchain

## Blockchain configuration (Hyperledger fabric)

- To build docker environment and create the channel execute:

```sh
./hyper.sh up createChannel
```

- To install smart contract execute:

```sh
./hyper.sh deployCC
```

- To interact with the blockchain network, we must set the following environment variables

```bash
export PATH=${PWD}/bin:$PATH
export FABRIC_CFG_PATH=$PWD/config/
```

- To work with the peer Org1 we must configure the following variables:

```bash
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```

- To work with the peer Org2 we must configure the following variables:

```bash
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```

## Apache Ranger configuration

- before executing any command, you must create the network that will interact the apache ranger and hadoop containers.

```bash
docker network create ranger-env
```

- To raise the apache ranger docker environment, we need to access the ranger folder and execute the command:

```bash
docker-compose up -d
```

- The following credentials are used to log in to the platform:

```bash
user: admin
pass: rangeradmin1
```
