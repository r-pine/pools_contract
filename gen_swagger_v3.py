import requests
import json

url = "https://converter.swagger.io/api/convert"

with open("gen/openapi/v2/pools.swagger-v2.json", "r") as f:
    data = json.load(f)

res = requests.post(url=url, json=data)

new_spec = res.json()

with open("gen/openapi/v3/pools.swagger-v3.json", "w") as f:
    json.dump(new_spec, f, indent=4, ensure_ascii=False)

print("gen spec v3 success!")