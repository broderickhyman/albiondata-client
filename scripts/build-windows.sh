#!/usr/bin/env bash

apt-get update && apt-get install -y build-essential mingw-w64
go get -u github.com/golang/dep/cmd/dep
dep ensure

export CGO_CPPFLAGS="-I $GOPATH/src/github.com/regner/albionmarket-client/thirdparty/WpdPack/Include/"
export CGO_LDFLAGS="-L $GOPATH/src/github.com/regner/albionmarket-client/thirdparty/WpdPack/Lib/x64/"
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1
export CXX=x86_64-w64-mingw32-g++
export CC=x86_64-w64-mingw32-gcc
go build -ldflags "-s -w -X main.version=$CIRCLE_TAG" -o albionmarket-client.exe -v -x cmd/albionmarket-client/albionmarket-client.go

gzip -9 albionmarket-client.exe
mv albionmarket-client.exe.gz update-windows-amd64.exe.gz
