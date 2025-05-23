#!/bin/bash

VERSIONTAG=1.0.0

# Build the Docker image.
docker build -t jwtsample:${VERSIONTAG} .

# Check if the image was built successfully.
docker images | grep jwtsample

# If you want to pus this image later to an image registry
# docker tag jwtsample:${VERSIONTAG} <Your image registry URL>/jwtsample:${VERSIONTAG}
# docker push <Your image registry URL>/jwtsample:${VERSIONTAG}

# Check the images
# curl -L https://<Your image registry URL>/v2/_catalog
# curl -L https://<Your image registry URL>/v2/jwtsample/tags/list
