#!/bin/bash -ex

DYNAMODB_ENDPOINT=${LAMBDA_ENDPOINT:-http://localhost:4569}
TABLE_NAME=${STREAM_NAME:-awsutil-test}

aws dynamodb create-table --endpoint-url ${DYNAMODB_ENDPOINT} \
    --table-name ${TABLE_NAME} \
    --key-schema '{"AttributeName": "key", "KeyType": "HASH"}' \
    --attribute-definitions '{"AttributeName": "key", "AttributeType": "S"}' \
    --provisioned-throughput '{"ReadCapacityUnits": 5, "WriteCapacityUnits": 5}'
