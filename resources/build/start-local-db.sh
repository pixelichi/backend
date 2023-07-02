#!/bin/bash -e

mkdir -p ~/mnt/local/minio

container_name=$1
port=$2
db_name=$3
db_user=$4
db_pass=$5

if [ "$(docker ps -aq -f name=$container_name)" ]; then
  docker container start $container_name
  printf "started!\n\n"
  exit 0
fi

docker run -d --name $container_name \
  -e POSTGRES_USER=$db_user \
  -e POSTGRES_PASSWORD=$db_pass \
  -e POSTGRES_DB=$db_name \
  -p $port:5432 postgres:15
printf "Started!\n\n"
