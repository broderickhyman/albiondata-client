#!/usr/bin/env bash

set -e

go run -ldflags="-w -s -X main.version=dev" cmd/albiondata-client/albiondata-client.go
