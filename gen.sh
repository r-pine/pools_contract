#!/bin/bash

echo "generate protobuf models and methods"

PROTO_DIR=$(pwd)

buf generate &&
echo $PROTO_DIR
for dir in $PROTO_DIR/gen/proto/*/; do \
  echo $dir && \
  cd $dir && \
  go mod init "github.com/r-pine/pools_contract/gen/proto/$(basename "$dir")" && \
  GOPRIVATE=github.com/r-pine/pools_contract go mod tidy; \
done

echo "successfully"
echo "generate swagger documentation"

cd ../../..
ls
sh ./gen_swagger.sh

echo "successfully"