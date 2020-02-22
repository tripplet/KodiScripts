#!/bin/sh
go generate
env GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o rpc-proxy
