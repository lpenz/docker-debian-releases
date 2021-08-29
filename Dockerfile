FROM debian:bullseye
MAINTAINER Leandro Lisboa Penz <lpenz@lpenz.org>

# install debian packages:
ENV DEBIAN_FRONTEND=noninteractive
RUN set -e -x; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        ca-certificates locales scons git ssh \
        golang golang-golang-x-net-dev

# setup su and locale
RUN set -e -x; \
    sed -i '/pam_rootok.so$/aauth sufficient pam_permit.so' /etc/pam.d/su; \
    echo 'en_US.UTF-8 UTF-8' >> /etc/locale.gen; locale-gen
ENV LC_ALL=en_US.UTF-8

# setup Debian GOPATH
ENV GOPATH=/usr/share/go:/usr/share/gocode XDG_CACHE_HOME=/tmp

CMD ["scons"]
