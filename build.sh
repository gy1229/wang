#!/bin/bash
echo "Kill old server"
kill -9 $(lsof -i :18888 | awk '{print $2}')

go build .

./oa