#!/usr/bin/env bash

set -eo pipefail

sudo apt-get install -y libpcap-dev
go get -u github.com/golang/dep/cmd/dep
dep ensure

go build -ldflags "-s -w -X main.version=$CIRCLE_TAG" cmd/albionmarket-client/albionmarket-client.go

gzip -9 albionmarket-client
mv albionmarket-client.gz update-linux-amd64.gz
