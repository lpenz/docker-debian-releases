#!/bin/bash

set -e -x

IFS=-
read -r MIRROR DIST ARCH VARIANT <<EOF
$TRAVIS_BRANCH
EOF
unset IFS

: This script converts travis environment vars into arguments for
: docker-create-debian-image
: MIRROR="$MIRROR" DIST="$DIST" ARCH="$ARCH" VARIANT="$VARIANT"

IMAGE="$DOCKERHUB_USERNAME/debian-${DIST}-${ARCH}"

ARGS=()

if [ "$ARCH" != "i386" ] && [ "$ARCH" != "amd64" ]; then
    ARGS+=("-q")
fi

if [ -n "$VARIANT" ]; then
    ARGS+=("-t" "$VARIANT")
    IMAGE+="-$VARIANT"
fi

./docker-create-debian-image -m "$MIRROR" -i "$IMAGE" "${ARGS[@]}" "$DIST" "$ARCH"
