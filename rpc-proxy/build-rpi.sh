#!/bin/sh
go generate
env GOOS=linux GOARCH=arm go build -o rpc-proxy
