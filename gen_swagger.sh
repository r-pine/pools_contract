#!/bin/bash

files=$(find gen  -name "*.json")

jq -s 'reduce .[] as $item ({}; . * $item)' $files > gen/openapi/v2/pools.swagger-v2.json

echo "generate swagger specification in >> gen/openapi/v2/pools.swagger.json"

rm -rf gen/openapi/v2/proto

mkdir gen/openapi/v3

python3 gen_swagger_v3.py