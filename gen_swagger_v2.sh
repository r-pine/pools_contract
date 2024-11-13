#!/bin/bash

files=$(find gen  -name "*.json")

jq -s 'reduce .[] as $item ({}; . * $item)' $files > gen/openapi/v2/gateway.swagger-v2.json

echo "generate swagger specification in >> gen/openapi/v2/gateway.swagger.json"

rm -rf gen/openapi/v2/proto