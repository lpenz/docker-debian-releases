FROM debian:bullseye
MAINTAINER Leandro Lisboa Penz <lpenz@lpenz.org>

# install debian packages:
ENV DEBIAN_FRONTEND=noninteractive
RUN set -e -x; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        ca-certificates locales scons git ssh \
        golang golang-golang-x-net-dev binutils \
        curl sudo debootstrap docker.io; \
    sed -i '/pam_rootok.so$/aauth sufficient pam_permit.so' /etc/pam.d/su; \
    echo 'en_US.UTF-8 UTF-8' >> /etc/locale.gen; locale-gen

COPY debootstrap_scripts/* /usr/share/debootstrap/scripts/

ENV LC_ALL=en_US.UTF-8 \
    GOPATH=/usr/share/go:/usr/share/gocode \
    XDG_CACHE_HOME=/tmp

CMD ["scons"]

# To use this dockerfile to build an image:
# docker build -t docker-debian-releases .
# docker run --rm -u 0 -v "${PWD}:$PWD" -w "$PWD" -v /var/run/docker.sock:/var/run/docker.sock \
#     docker-debian-releases ./docker-create-debian-image -m devuan chimaera i386
