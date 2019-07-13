#!/bin/bash

set -e -x

export TRAVIS_BRANCH=${1?usage: "$0" mirror-dist-arch}

./travis-script.sh _image.txt
# shellcheck disable=SC1091
. ./_image.txt
export IMAGE DESCRIPTION
./travis-deploy.sh

