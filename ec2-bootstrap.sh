#!/bin/bash

#
# This script bootstrap an EC2 instance
# that runs a Dockerized version of the web server
# 
# Author: nguyen.tran

# Environment variables
PORT=3000
DOCKERHUB_USERNAME=
DOCKERHUB_PASSWORD=
DOCKERHUB_IMAGE_NAME=
MONGODB_USERNAME=
MONGODB_PASSWORD=
MONGODB_DATABASE=

# General updates
apt -y update

# Install Docker
apt -y install apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --yes --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
apt -y update
apt-cache -y policy docker-ce
apt -y install docker-ce
systemctl start docker

# Pull the necessary image
docker login -u $DOCKERHUB_USERNAME -p $DOCKERHUB_PASSWORD
docker pull $DOCKERHUB_IMAGE_NAME

# Runs the image
docker run -it -p $PORT:$PORT \
    -e GIN_MODE=release \
    -e PORT=$PORT \
    -e LOG_LEVEL=INFO \
    -e MONGO_USERNAME=$MONGODB_USERNAME \
    -e MONGO_PASSWORD=$MONGODB_PASSWORD \
    -e MONGO_CLUSTER_ID=$MONGODB_DATABASE \
    $DOCKERHUB_IMAGE_NAME 
