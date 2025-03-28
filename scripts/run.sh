#!/usr/bin/env sh

set -e

docker run --rm --net=host -v $(pwd)/config.toml:/workspace/config.toml -v ./logs:/workspace/logs 117503445/guardix