#!/usr/bin/env bash

set -eo pipefail

apt-get update && apt-get install -y libpcap-dev
go get -u github.com/golang/dep/cmd/dep
dep ensure

export OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1
export MACOSX_DEPLOYMENT_TARGET=10.6
export CC=/usr/osxcross/bin/o64-clang
export CXX=/usr/osxcross/bin/o64-clang++
export GOOS=darwin
export GOARCH=amd64 CGO_ENABLED=1
go build -ldflags "-s -w -X main.version=$CIRCLE_TAG" cmd/albiondata-client/albiondata-client.go

gzip -9 albiondata-client
mv albiondata-client.gz update-darwin-amd64.gz