#!/usr/bin/env bash

protoDir="."
outDir="./user"

protoc --proto_path=${protoDir} \
  --php_out=${outDir} \
  --grpc_out=${outDir} \
  --plugin=protoc-gen-grpc=/Users/liubing/go/src/staff_go/grpc/php-client/grpc/bins/opt/grpc_php_plugin \
  ${protoDir}/*.proto