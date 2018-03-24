#!/usr/bin/env bash
# docker login --username=shotcounterapp --email=shotcounterapp@gmail.com

# This should be in a docker image probably
cd go-workspace && export GOPATH=$PWD && cd src/github.com/counting-frontend && go test ./... && cd ../../../..

docker build --no-cache -t shot-counter-frontend -f Dockerfile.build .

docker tag shot-counter-frontend shotcounterapp/shot-counter-frontend
docker push shotcounterapp/shot-counter-frontend
