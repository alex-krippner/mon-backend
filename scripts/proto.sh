#!/bin/bash
set -e

readonly service=$1

# Get the directory of the script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

PROTO_PATH="$SCRIPT_DIR/../api/protobuf"
GEN_PROTO_DIR="$SCRIPT_DIR/../adapters/genproto"

# Check if the file exists
if [ ! -f "$PROTO_PATH/$service.proto" ]; then
    echo "Error: Protobuf file not found at $PROTO_PATH"
    exit 1
fi

protoc \
  --proto_path=$PROTO_PATH "$PROTO_PATH/$service.proto" \
  "--go_out=$GEN_PROTO_DIR/$service" --go_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  "--go-grpc_out=$GEN_PROTO_DIR/$service" --go-grpc_opt=paths=source_relative