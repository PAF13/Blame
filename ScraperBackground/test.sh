#!/usr/bin/env sh

set -e

echo "=== Test GET localhost:8080 ==="
curl -X GET -s localhost:8080 | rg -q "get" 

echo "=== Test POST localhost:8080 ==="
curl -X POST -s localhost:8080 | rg -q "post"

echo "=== Insert Test Data 1 ==="

curl -X POST localhost:8080 -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

echo "=== Insert Test Data 2 ==="

curl -X POST localhost:8080 -d \
'{"activity": {"description": "cross country skiing is horrible and cold", "time":"2021-12-09T16:56:12Z"}}'

echo "=== Insert Test Data 3 ==="

curl -X POST localhost:8080 -d \
'{"activity": {"description": "sledding with nephew", "time":"2021-12-09T16:56:23Z"}}'


echo "=== Test Descriptions ==="

curl -X GET localhost:8080 -d '{"id": 0}' | grep -q 'christmas eve bike class'
curl -X GET localhost:8080 -d '{"id": 1}' | grep -q 'cross country skiing'
curl -X GET localhost:8080 -d '{"id": 2}'

echo "Success"

