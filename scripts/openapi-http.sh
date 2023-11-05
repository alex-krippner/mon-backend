#!/bin/bash
set -e

readonly output_dir="$1"
readonly package="$2"

oapi-codegen -generate types -o "$output_dir/openapi_client_gen.go" -package "$package" "api/mon.yml"
oapi-codegen -generate chi-server -o "$output_dir/openapi_api.gen.go" -package "$package" "api/mon.yml"