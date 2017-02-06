@echo off

mkdir .gopath
mkdir .gopath\windows

SET GOPATH=%CD%\.gopath\windows
SET PATH=%MINGW64_PATH%\bin;%PATH%

echo ############################
echo ### get libraries
echo ############################
echo GOROOT=%GOROOT%
echo GOPATH=%GOPATH%
go get github.com/stretchr/testify/assert
go get github.com/urfave/cli

echo ############################
echo ### prebuilt cgo repo
echo ############################
go install github.com/stretchr/testify/assert
go install github.com/urfave/cli

echo ############################
echo ### build
echo ############################
go build src\main.go
