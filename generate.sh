#!/bin/bash

# Check if protoc is installed
if ! [ -x "$(command -v protoc)" ]; then
  echo 'Error: protoc is not installed.' >&2
  exit 1
fi

# Generate Go code for Protobuf messages
protoc -I protos --go_out=./go protos/*.proto

# Generate Go code for gRPC service definitions
protoc -I protos --go-grpc_out=./go protos/*.proto
