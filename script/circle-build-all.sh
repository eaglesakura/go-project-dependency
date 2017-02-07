#! /bin/sh -eu

rm -rf ./ci-release
mkdir ./ci-release

echo "########################"
echo "## Build for Linux"
echo "########################"
export GOOS=linux
export GOARCH=amd64

go build -o prjdep
mkdir ./ci-release/linux
mv ./prjdep ./ci-release/linux

echo "########################"
echo "## Build for Mac"
echo "########################"
export GOOS=darwin
export GOARCH=amd64

go build -o prjdep
mkdir ./ci-release/mac
mv ./prjdep ./ci-release/mac

echo "########################"
echo "## Build for Windows"
echo "########################"
export GOOS=windows
export GOARCH=amd64

go build -o prjdep.exe
mkdir ./ci-release/windows
mv ./prjdep.exe ./ci-release/windows
