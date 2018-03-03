#!/bin/bash

export GOPATH=/app/go-workspace && export GOBIN=/app && go install /app/go-workspace/src/github.com/counting-frontend/cmd/counting-frontend/main.go
