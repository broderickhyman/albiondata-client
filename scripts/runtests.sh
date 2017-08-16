#!/usr/bin/env bash

set -e

echo $CIRCLE_TAG
go run cmd/albionmarket-client/albionmarket-client.go -o fixtures/refresh-market-02.pcap
