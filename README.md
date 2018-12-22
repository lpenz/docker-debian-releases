[![Build Status](https://travis-ci.com/lpenz/docker-debian-releases.png?branch=master)](https://travis-ci.com/lpenz/docker-debian-releases)

docker-debian-releases
======================

This repository has a single that can be used to create docker images of
historic Debian releases.

The travis-ci build matrix defined here updates the images at
https://hub.docker.com/r/lpenz/ with all releases of Debian since Potato, for
all architectures that *debootstrap* can build with qemu. The status of the
jobs can be seen here: https://travis-ci.com/lpenz/docker-debian-releases

| Release       | Architecture | Variant | Status                                                                                                                                                               | 
| -------       | ------------ | ------- | ------                                                                                                                                                               | 
| 2.2 (Potato)  | alpha        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-potato-alpha)            | 
| 2.2 (Potato)  | arm          |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-potato-arm)              | 
| 2.2 (Potato)  | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-potato-i386)             | 
| 3.0 (Woody)   | alpha        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-woody-alpha)             | 
| 3.0 (Woody)   | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-woody-i386)              | 
| 3.1 (Sarge)   | alpha        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-sarge-alpha)             | 
| 3.1 (Sarge)   | arm          |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-sarge-arm)               | 
| 3.1 (Sarge)   | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-sarge-i386)              | 
| 3.1 (Sarge)   | mips         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-sarge-mips)              | 
| 3.1 (Sarge)   | mipsel       |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-sarge-mipsel)            | 
| 3.1 (Sarge)   | powerpc      |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-sarge-powerpc)           | 
| 4.0 (Etch)    | alpha        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-alpha)              | 
| 4.0 (Etch)    | alpha        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-alpha-minbase)      | 
| 4.0 (Etch)    | amd64        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-amd64)              | 
| 4.0 (Etch)    | amd64        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-amd64-minbase)      | 
| 4.0 (Etch)    | arm          |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-arm)                | 
| 4.0 (Etch)    | arm          | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-arm-minbase)        | 
| 4.0 (Etch)    | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-i386)               | 
| 4.0 (Etch)    | i386         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-i386-minbase)       | 
| 4.0 (Etch)    | mips         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-mips)               | 
| 4.0 (Etch)    | mipsel       |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-mipsel)             | 
| 4.0 (Etch)    | mipsel       | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-mipsel-minbase)     | 
| 4.0 (Etch)    | mips         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-mips-minbase)       | 
| 4.0 (Etch)    | powerpc      |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-powerpc)            | 
| 4.0 (Etch)    | powerpc      | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-etch-powerpc-minbase)    | 
| 5.0 (Lenny)   | alpha        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-alpha)             | 
| 5.0 (Lenny)   | alpha        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-alpha-minbase)     | 
| 5.0 (Lenny)   | amd64        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-amd64)             | 
| 5.0 (Lenny)   | amd64        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-amd64-minbase)     | 
| 5.0 (Lenny)   | arm          |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-arm)               | 
| 5.0 (Lenny)   | armel        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-armel)             | 
| 5.0 (Lenny)   | armel        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-armel-minbase)     | 
| 5.0 (Lenny)   | arm          | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-arm-minbase)       | 
| 5.0 (Lenny)   | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-i386)              | 
| 5.0 (Lenny)   | i386         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-i386-minbase)      | 
| 5.0 (Lenny)   | mips         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-mips)              | 
| 5.0 (Lenny)   | mipsel       |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-mipsel)            | 
| 5.0 (Lenny)   | mipsel       | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-mipsel-minbase)    | 
| 5.0 (Lenny)   | mips         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-mips-minbase)      | 
| 5.0 (Lenny)   | powerpc      |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-powerpc)           | 
| 5.0 (Lenny)   | powerpc      | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-lenny-powerpc-minbase)   | 
| 6.0 (Squeeze) | amd64        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-amd64)           | 
| 6.0 (Squeeze) | amd64        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-amd64-minbase)   | 
| 6.0 (Squeeze) | armel        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-armel)           | 
| 6.0 (Squeeze) | armel        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-armel-minbase)   | 
| 6.0 (Squeeze) | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-i386)            | 
| 6.0 (Squeeze) | i386         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-i386-minbase)    | 
| 6.0 (Squeeze) | mips         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-mips)            | 
| 6.0 (Squeeze) | mipsel       |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-mipsel)          | 
| 6.0 (Squeeze) | mipsel       | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-mipsel-minbase)  | 
| 6.0 (Squeeze) | mips         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-mips-minbase)    | 
| 6.0 (Squeeze) | powerpc      |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-powerpc)         | 
| 6.0 (Squeeze) | powerpc      | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-squeeze-powerpc-minbase) | 
| 7 (Wheezy)    | amd64        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-amd64)            | 
| 7 (Wheezy)    | amd64        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-amd64-minbase)    | 
| 7 (Wheezy)    | armel        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-armel)            | 
| 7 (Wheezy)    | armel        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-armel-minbase)    | 
| 7 (Wheezy)    | armhf        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-armhf)            | 
| 7 (Wheezy)    | armhf        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-armhf-minbase)    | 
| 7 (Wheezy)    | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-i386)             | 
| 7 (Wheezy)    | i386         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-i386-minbase)     | 
| 7 (Wheezy)    | mips         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-mips)             | 
| 7 (Wheezy)    | mipsel       |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-mipsel)           | 
| 7 (Wheezy)    | mipsel       | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-mipsel-minbase)   | 
| 7 (Wheezy)    | mips         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-mips-minbase)     | 
| 7 (Wheezy)    | powerpc      |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-powerpc)          | 
| 7 (Wheezy)    | powerpc      | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-powerpc-minbase)  | 
| 7 (Wheezy)    | s390x        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-s390x)            | 
| 7 (Wheezy)    | s390x        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-wheezy-s390x-minbase)    | 
| 8 (Jessie)    | amd64        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-amd64)            | 
| 8 (Jessie)    | amd64        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-amd64-minbase)    | 
| 8 (Jessie)    | arm64        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-arm64)            | 
| 8 (Jessie)    | arm64        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-arm64-minbase)    | 
| 8 (Jessie)    | armel        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-armel)            | 
| 8 (Jessie)    | armel        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-armel-minbase)    | 
| 8 (Jessie)    | armhf        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-armhf)            | 
| 8 (Jessie)    | armhf        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-armhf-minbase)    | 
| 8 (Jessie)    | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-i386)             | 
| 8 (Jessie)    | i386         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-i386-minbase)     | 
| 8 (Jessie)    | mips         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-mips)             | 
| 8 (Jessie)    | mipsel       |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-mipsel)           | 
| 8 (Jessie)    | mipsel       | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-mipsel-minbase)   | 
| 8 (Jessie)    | mips         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-mips-minbase)     | 
| 8 (Jessie)    | powerpc      |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-powerpc)          | 
| 8 (Jessie)    | powerpc      | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-powerpc-minbase)  | 
| 8 (Jessie)    | ppc64el      |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-ppc64el)          | 
| 8 (Jessie)    | ppc64el      | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-ppc64el-minbase)  | 
| 8 (Jessie)    | s390x        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-s390x)            | 
| 8 (Jessie)    | s390x        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-jessie-s390x-minbase)    | 
| 9 (Stretch)   | amd64        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-amd64)           | 
| 9 (Stretch)   | amd64        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-amd64-minbase)   | 
| 9 (Stretch)   | arm64        |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-arm64)           | 
| 9 (Stretch)   | arm64        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-arm64-minbase)   | 
| 9 (Stretch)   | armel        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-armel-minbase)   | 
| 9 (Stretch)   | armhf        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-armhf-minbase)   | 
| 9 (Stretch)   | i386         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-i386)            | 
| 9 (Stretch)   | i386         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-i386-minbase)    | 
| 9 (Stretch)   | mips         |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-mips)            | 
| 9 (Stretch)   | mipsel       |         | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-mipsel)          | 
| 9 (Stretch)   | mipsel       | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-mipsel-minbase)  | 
| 9 (Stretch)   | mips         | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-mips-minbase)    | 
| 9 (Stretch)   | s390x        | minbase | [![ok](https://img.shields.io/badge/docker-ok-brightgreen.svg?logo=docker)](https://cloud.docker.com/u/lpenz/repository/docker/lpenz/debian-stretch-s390x-minbase)   | 


Some images couldn't be built, mostly due to travis timeouts and lack of support in qemu:

| Release       | Architecture | Status                                                                                                                                       | 
| -------       | ------------ | ------                                                                                                                                       | 
| 2.2 (Potato)  | m68k         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465551) | 
| 2.2 (Potato)  | powerpc      | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465552) | 
| 2.2 (Potato)  | sparc        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465553) | 
| 3.0 (Woody)   | arm          | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465555) | 
| 3.0 (Woody)   | hppa         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465556) | 
| 3.0 (Woody)   | ia64         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465558) | 
| 3.0 (Woody)   | m68k         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465559) | 
| 3.0 (Woody)   | mips         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465560) | 
| 3.0 (Woody)   | mipsel       | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465561) | 
| 3.0 (Woody)   | powerpc      | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465562) | 
| 3.0 (Woody)   | s390         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465563) | 
| 3.0 (Woody)   | sparc        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465564) | 
| 3.1 (Sarge)   | hppa         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465567) | 
| 3.1 (Sarge)   | ia64         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465570) | 
| 3.1 (Sarge)   | m68k         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465571) | 
| 3.1 (Sarge)   | s390         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465575) | 
| 3.1 (Sarge)   | sparc        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465576) | 
| 4.0 (Etch)    | hppa         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465583) | 
| 4.0 (Etch)    | ia64         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465587) | 
| 4.0 (Etch)    | s390         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465595) | 
| 4.0 (Etch)    | sparc        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465597) | 
| 5.0 (Lenny)   | hppa         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465612) | 
| 5.0 (Lenny)   | ia64         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465617) | 
| 5.0 (Lenny)   | s390         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465625) | 
| 5.0 (Lenny)   | sparc        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465627) | 
| 6.0 (Squeeze) | ia64         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465639) | 
| 6.0 (Squeeze) | s390         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465651) | 
| 6.0 (Squeeze) | sparc        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465654) | 
| 7 (Wheezy)    | ia64         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465664) | 
| 7 (Wheezy)    | s390         | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465677) | 
| 7 (Wheezy)    | sparc        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465681) | 
| 9 (Stretch)   | armel        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465708) | 
| 9 (Stretch)   | armhf        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465712) | 
| 9 (Stretch)   | mips64el     | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465718) | 
| 9 (Stretch)   | ppc64el      | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465723) | 
| 9 (Stretch)   | s390x        | [![error](https://img.shields.io/badge/travis-error-red.svg?logo=travis)](https://travis-ci.com/lpenz/docker-debian-releases/jobs/164465725) | 

