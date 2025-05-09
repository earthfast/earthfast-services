#!/bin/bash

# this script is used to generate a docker-compose.yml file with the specified number of content nodes

rm docker-compose.yml
cat docker-compose.base.yml > docker-compose.yml

NUM_CONTENT_NODES=$1
echo "Generating $NUM_CONTENT_NODES content nodes"

for i in $(seq 0 $(($NUM_CONTENT_NODES-1)))
do
  port=$((30083 + i))
  container_id=$(($i+1))
  echo "
  content-$container_id:
    extends:
      service: content-template
    environment:
      - NODE_ID=0x405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3b$port
      # - DISABLE_NGINX=true
    ports:
      - $port:80
  " >> docker-compose.yml
done
