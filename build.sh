#!/bin/bash

rm -r ./bin
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8 ./build_docker.sh

docker build -t tzapil/gomal:v0.5 -f Dockerfile .