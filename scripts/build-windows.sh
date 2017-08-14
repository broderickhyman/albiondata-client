#!/usr/bin/env bash

export CGO_CPPFLAGS="-I $GOPATH/src/github.com/regner/albionmarket-client/thirdparty/WpdPack/Include/"
export CGO_LDFLAGS="-L $GOPATH/src/github.com/regner/albionmarket-client/thirdparty/WpdPack/Lib/x64/"
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1
export CXX=x86_64-w64-mingw32-g++
export CC=x86_64-w64-mingw32-gcc
go build -ldflags "-s -w -X main.version=`git describe --tags`" -o albionmarket-client.exe -v -x cmd/albionmarket-client/albionmarket-client.go
