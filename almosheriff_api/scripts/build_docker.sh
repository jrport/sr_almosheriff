#! /usr/bin/bash

# Build the image
sudo docker buildx build . --tag postgres_image

# Run the container
sudo docker run --name postgres_container -d -p 5000:5432 postgres_image
