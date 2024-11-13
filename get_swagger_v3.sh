#!/bin/bash

mkdir -p "gen/openapi/v3/"

files=$(find proto  -name "*.proto")

protoc --openapi_out=gen/openapi/v3/ --openapi_opt=version=1.0.0 --openapi_opt=title="R-pine API" $files
sed -i 's|/v1||g' gen/openapi/v3/openapi.yaml
sed -i 's|info:|servers:\n  - description: API server version 1\n    url: https://api.rpine.xyz/v1\ninfo:|g' gen/openapi/v3/openapi.yaml
#- description: API server version 2\n    url: https://api.rpine.xyz/v2\n  - description: API server version 3\n    url: https://api.rpine.xyz/v3\n
python3 add_jwt_openapi.py gen/openapi/v3/openapi.yaml