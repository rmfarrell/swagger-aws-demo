yaml2json api.yml > api.json

curl -k -X POST "https://apigateway.us-east-1.amazonaws.com/"

curl -i \
    https://apigateway.us-east-1.amazonaws.com/restapis/test?mode=import \
    -H "Accept: application/yaml" \
    -F "file=api.yml" \
    -X POST -d '{"Content-Type": "application/json"}'

