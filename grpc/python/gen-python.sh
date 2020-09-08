#!/usr/bin/env bash

protoDir="."
outDir="./user"

/usr/local/Cellar/python3/3.7.4_1/bin/python3.7 -m grpc_tools.protoc -I ${protoDir}/ --python_out=${outDir} --grpc_python_out=${outDir} ${protoDir}/*proto