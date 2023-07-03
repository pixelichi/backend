#!/bin/bash -e

mkdir -p ~/mnt/local/minio

container_name=$1
access_key=$2
secret=$3
public_minio_url=$4

if [ "$(docker ps -aq -f name=$container_name)" ]; then
    docker container start $container_name
    printf "started!\n\n"
    exit 0
fi

docker run -d \
    -p 7777:9000 \
    -p 7778:9090 \
    --name $container_name \
    -v ~/mnt/local/minio/data:/data \
    -e "MINIO_ACCESS_KEY=$access_key" \
    -e "MINIO_SECRET_KEY=$secret" \
    quay.io/minio/minio server /data --console-address ":9090"

printf "Started!\n\n"
