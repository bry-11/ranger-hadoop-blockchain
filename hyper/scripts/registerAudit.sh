#!/bin/bash

# Asignar los parámetros a variables
mode=$1
jsonData=$2


function initializeAudits() {
    peer chaincode invoke \
        -o orderer.example.com:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls \
        --cafile ${PWD}/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
        -C mychannel \
        -n basic \
        --peerAddresses peer0.org1.example.com:7051 \
        --tlsRootCertFiles ${PWD}/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
        --peerAddresses peer0.org2.example.com:9051 \
        --tlsRootCertFiles ${PWD}/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
        -c '{"function":"InitLedger","Args":[]}'
}

function getAllAudits() {
    peer chaincode query \
        -C mychannel \
        -n basic \
        -c '{"function":"GetAllAudits","Args":[]}'
}

function registerAudit() {
    peer chaincode invoke \
        -o orderer.example.com:7050 \
        --ordererTLSHostnameOverride orderer.example.com \
        --tls \
        --cafile ${PWD}/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
        -C mychannel \
        -n basic \
        --peerAddresses peer0.org1.example.com:7051 \
        --tlsRootCertFiles ${PWD}/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
        --peerAddresses peer0.org2.example.com:9051 \
        --tlsRootCertFiles ${PWD}/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
        -c "{\"function\":\"CreateAudit\",\"Args\":[\"$uuid\",\"$jsonData\"]}"
}

if [ "$mode" == "init" ]; then
    initializeAudits
elif [ "$mode" == "getAll" ]; then
    getAllAudits
elif [ "$mode" == "register" ]; then
    uuid=$(openssl rand -hex 16 | sed 's/^\(.\{8\}\)\(.\{4\}\)\(.\{4\}\)\(.\{4\}\)\(.\{12\}\)$/\1-\2-\3-\4-\5/')
    registerAudit
    # Mostrar el UUID generado
    echo "Registro enviado con UUID: $uuid"
else
    exit 1
fi
