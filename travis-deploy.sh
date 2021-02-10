#!/bin/bash

if [ -z "$IMAGE" ]; then
    echo Unspecified IMAGE >&2
    exit 1
fi

if [ -z "$DESCRIPTION" ]; then
    echo Unspecified DESCRIPTION >&2
    exit 1
fi

set -e -x

docker push "$IMAGE"
./dockerhub-set-descriptions "$IMAGE" "$DESCRIPTION" "Automatically updated via https://github.com/lpenz/docker-debian-releases/"

