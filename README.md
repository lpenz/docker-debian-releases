[![Build Status for master](https://img.shields.io/travis/com/lpenz/docker-debian-releases/master.svg?label=master)](https://travis-ci.com/lpenz/docker-debian-releases)
[![Build Status for debian](https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian.svg?label=debian)](https://travis-ci.com/lpenz/docker-debian-releases)
[![Build Status for ubuntu](https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu.svg?label=ubuntu)](https://travis-ci.com/lpenz/docker-debian-releases)

docker-debian-releases
======================

This repository has a script that creates docker images of
Debian-based system using debootstrap, for various architectures.

It's also hooked up in travis with a build matrix that creates
historical Debian and Ubuntu release in the *debian* and *ubuntu*
branches. You can check out the docker images at
https://hub.docker.com/r/lpenz/ - all releases of Debian since Potato
are there, for various architectures. The status of the jobs can be
seen here: https://travis-ci.com/lpenz/docker-debian-releases
