#!/usr/bin/env sh

set -e

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .
docker build -t 117503445/guardix .