#!/usr/bin/env bash

set -eo pipefail

sudo apt-get install -y libpcap-dev

go build -ldflags "-s -w -X main.version=$CIRCLE_TAG" albiondata-client.go

gzip -9 albiondata-client
mv albiondata-client.gz update-linux-amd64.gz
