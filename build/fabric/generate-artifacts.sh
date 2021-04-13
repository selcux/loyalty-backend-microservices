#!/usr/bin/env bash

function generateCerts(){
    echo
    echo "##########################################################"
    echo "##### Generate certificates using cryptogen tool #########"
    echo "##########################################################"
    if [ -d ./crypto-config ]; then
            rm -rf ./crypto-config
    fi

    cryptogen generate --config=./crypto-config.yaml
    echo
}


function generateChannelArtifacts(){

    if [ ! -d ./channel-artifacts ]; then
                mkdir channel-artifacts
    fi

    echo
    echo "#################################################################"
    echo "### Generating channel configuration transaction 'channel.tx' ###"
    echo "#################################################################"

    configtxgen -profile TwoOrgsOrdererGenesis -channelID testchainid -outputBlock ./channel-artifacts/genesis.block
    configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID "mychannel"


    echo
    echo "#################################################################"
    echo "#######    Generating anchor peer update for Org1MSP   ##########"
    echo "#################################################################"
    configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/org1Anchors.tx  -channelID "mychannel" -asOrg org1MSP

    echo
    echo "#################################################################"
    echo "#######    Generating anchor peer update for Org2MSP   ##########"
    echo "#################################################################"
    configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/org2Anchors.tx  -channelID "mychannel" -asOrg org2MSP
    echo
}

generateCerts
generateChannelArtifacts
