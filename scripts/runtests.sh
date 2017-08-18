#!/usr/bin/env bash

set -e

go run cmd/albiondata-client/albiondata-client.go -o fixtures/refresh-market-02.pcap
