#!/bin/bash -ex

PLUGIN=wallet-plugin

mkdir -p api/server
go generate

# Build the binary
mkdir -p "build/$PLUGIN/"
CGO_ENABLED="1" go build -o build/$PLUGIN/wallet-plugin cmd/massa-wallet/thyra-plugin-wallet.go
