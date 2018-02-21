#!/usr/bin/env bash

set -eo pipefail

#apt-get update && apt-get install -y libpcap-dev
#go get -u github.com/golang/dep/cmd/dep
#dep ensure

export OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1
xgo -ldflags="-s -w -X main.version=$CIRCLE_TAG" --targets=darwin-10.6/amd64 ./cmd/albiondata-client/

TEMP="albiondata-client"

rm -rf ./scripts/$TEMP
rm -rf ./scripts/update-darwin-amd64.zip
mkdir ./scripts/$TEMP
mv albiondata-client-darwin-10.6-amd64 ./scripts/$TEMP/albiondata-client-executable
cd scripts
cp run.command ./$TEMP/run.command
sudo chown -R ${USER}:${USER} ./$TEMP
chmod a+x ./$TEMP/*
zip update-darwin-amd64.zip -r ./$TEMP
