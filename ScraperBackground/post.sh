#!/usr/bin/env sh

set -e

echo "=== POST ==="

curl -X POST http://localhost:8080/tom -d '{"name":"Tom","age":21}'
	 
echo "Success"

