#!/usr/bin/env bash

set -eo pipefail

apt-get update && apt-get install -y libpcap-dev

export OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1
export MACOSX_DEPLOYMENT_TARGET=10.6
export CC=/usr/osxcross/bin/o64-clang
export CXX=/usr/osxcross/bin/o64-clang++
export GOOS=darwin
export GOARCH=amd64 CGO_ENABLED=1
go build -ldflags "-s -w -X main.version=$CIRCLE_TAG" albiondata-client.go


gzip -k9 albiondata-client
mv albiondata-client.gz update-darwin-amd64.gz


# Creates a zipped folder with a run.command file that runs the client under sudo
TEMP="albiondata-client"
ZIPNAME="albiondata-client-amd64-mac.zip"
rm -rfv ./scripts/$TEMP
rm -rfv ./$ZIPNAME
rm -rfv ./scripts/update-darwin-amd64.zip
mkdir -v ./scripts/$TEMP
cp -v albiondata-client ./scripts/$TEMP/albiondata-client-executable
cd scripts
cp -v run.command ./$TEMP/run.command
sudo chown -Rv ${USER}:${USER} ./$TEMP
chmod -v 777 ./$TEMP/*
zip -v ../$ZIPNAME -r ./"$TEMP"

# In theory the following works to create an app but there was a permissions issue when opening on the mac
# APP_NAME="Albion Data Client"
# TEMP="$APP_NAME".app
# ZIPNAME="albiondata-client-amd64-mac.zip"

# rm -rfv ./scripts/"$TEMP"
# rm -rfv ./scripts/"$ZIPNAME"
# mkdir -pv ./scripts/"$TEMP"/Contents/MacOS
# cp -v albiondata-client-darwin-10.6-amd64 ./scripts/"$TEMP"/Contents/MacOS/"$APP_NAME"
# sudo chown -Rv ${USER}:${USER} ./scripts/"$TEMP"
# chmod -v 777 ./scripts/"$TEMP"/*

# cd scripts
# zip -v ../$ZIPNAME -r ./"$TEMP"
