#!/bin/sh
env GOOS=linux GOARCH=arm go build -o rpc-proxy
