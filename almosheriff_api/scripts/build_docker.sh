#! /usr/bin/bash

# Build the image
docker buildx build . --tag postgres_image

# Build the container
docker run --name postgres_container -v ./docker/data:/var/lib/postgresql/data -d -p 5000:5432 postgres_image

# Start container
docker start postgres_container
