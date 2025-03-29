#!/usr/bin/env sh

set -e


docker build -t 117503445/guardix-builder -f ./scripts/docker/builder.dockerfile .
docker run --rm -v $(pwd):/workspace 117503445/guardix-builder