#!/bin/sh

PROJECT_ID=kyle-server-402706
REPOSITORY=kyle-registry
LOCATION=me-west1
IMAGE=chat-app-server
TAG=0.0.1
TAG_LATEST=latest

# dir
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
echo $SCRIPT_DIR
DOCKERFILE_PATH=$SCRIPT_DIR/server/Dockerfile
CONTEXT_PATH=$SCRIPT_DIR/server

set -e

# delete all images with tags (this will automatically handle dependencies)
gcloud artifacts docker images delete $LOCATION-docker.pkg.dev/$PROJECT_ID/$REPOSITORY/$IMAGE --delete-tags --quiet || true

# IMAGE_TAG=me-west1-docker.pkg.dev/kyle-server-402706/kyle-registry/chat-app-server:0.0.1
IMAGE_TAG=$LOCATION-docker.pkg.dev/$PROJECT_ID/$REPOSITORY/$IMAGE:$TAG
# --no-cache
docker buildx build --no-cache --build-arg=PROGRAM_VER=0.0.1 --load -t $IMAGE_TAG -f $DOCKERFILE_PATH $CONTEXT_PATH
# docker buildx build --no-cache --platform linux/amd64 --build-arg=PROGRAM_VER=0.0.1 --push -t $IMAGE_TAG -f $DOCKERFILE_PATH $CONTEXT_PATH
docker push $IMAGE_TAG

IMAGE_TAG_LATEST=$LOCATION-docker.pkg.dev/$PROJECT_ID/$REPOSITORY/$IMAGE:$TAG_LATEST
docker tag $IMAGE_TAG $IMAGE_TAG_LATEST
docker push $IMAGE_TAG_LATEST
