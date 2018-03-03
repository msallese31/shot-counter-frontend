#!/usr/bin/env bash
# docker login --username=shotcounterapp --email=shotcounterapp@gmail.com

docker build --no-cache -t shot-counter-frontend -f Dockerfile.build .

docker tag shot-counter-frontend shotcounterapp/shot-counter-frontend
docker push shotcounterapp/shot-counter-frontend