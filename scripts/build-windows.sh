#!/usr/bin/env bash

set -eo pipefail

# go get -u github.com/golang/dep/cmd/dep
# dep ensure

export CGO_CPPFLAGS="-I $GOPATH/src/github.com/broderickhyman/albiondata-client/thirdparty/WpdPack/Include/"
export CGO_LDFLAGS="-L $GOPATH/src/github.com/broderickhyman/albiondata-client/thirdparty/WpdPack/Lib/x64/"
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1
export CXX=x86_64-w64-mingw32-g++
export CC=x86_64-w64-mingw32-gcc
go build -ldflags "-s -w -X main.version=$CIRCLE_TAG" -o albiondata-client.exe -v -x cmd/albiondata-client/albiondata-client.go

# Add icon to the .exe
wine thirdparty/rcedit/rcedit.exe albiondata-client.exe --set-icon icon/albiondata-client.ico

# Make the NSIS Installer
cd pkg/nsis
make nsis
cd ../..

gzip -9 albiondata-client.exe
mv albiondata-client.exe.gz update-windows-amd64.exe.gz
