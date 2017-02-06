#!/bin/sh

mkdir .gopath
mkdir .gopath/unix

export GOPATH=$PWD/.gopath/unix
export CGO_ENABLED=0
export PROJECT_BUILD_TAGS=""

echo "############################"
echo "### get libraries"
echo "############################"
echo "GOROOT=${GOROOT}"
echo "GOPATH=${GOPATH}"
go get github.com/stretchr/testify/assert
go get github.com/urfave/cli
