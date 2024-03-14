#! /usr/bin/bash

sudo docker stop postgres_container
sudo docker rm postgres_container
sudo docker rmi postgres_image
