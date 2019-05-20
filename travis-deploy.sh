#!/bin/bash

IMAGE="$1"
DESCRIPTION="$2"

set -e -x

docker push "$IMAGE"
./dockerhub-set-descriptions "$IMAGE" "$DESCRIPTION" "Automatically updated via https://github.com/lpenz/docker-debian-releases/"
