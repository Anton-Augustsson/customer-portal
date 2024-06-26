#!/bin/bash

#export PATH=$PATH:$(go env GOPATH)/bin
export PATH=$PATH:/path/to/protoc-gen-go/directory
GOOGLEAPIS_DIR=../../googleapis

# $3 GOOGLEAPIS_DIR
grpc_update () {
    protoc --go_out=$1-service --go_opt=paths=source_relative \
        --go-grpc_out=$1-service --go-grpc_opt=paths=source_relative \
        api/$2-service/$2-service.proto
}

grpc_update robot-remote-controller robot-remote-controller
grpc_update robot-remote-controller subscription
grpc_update subscription subscription
