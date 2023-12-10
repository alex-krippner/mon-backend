#!/bin/bash
set -e

readonly service=$1

protoc \
  --proto_path=../api/protobuf "../api/protobuf/$service.proto" \
  "--go_out=../adapters/genproto/$service" --go_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  "--go-grpc_out=../adapters/genproto/$service" --go-grpc_opt=paths=source_relative