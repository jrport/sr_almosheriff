#! /usr/bin/bash

docker stop postgres_container
docker rm postgres_container
docker rmi postgres_image
