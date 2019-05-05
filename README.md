[![Build Status](https://img.shields.io/travis/com/lpenz/docker-debian-releases/master.svg?label=master)](https://travis-ci.com/lpenz/docker-debian-releases)


docker-debian-releases
======================

This repository has a script that creates docker images of
Debian-based system using debootstrap, for various architectures.

It's also hooked up in travis with a build matrix that creates
historical Debian and Ubuntu release in appropriately-named
branches. You can check out the docker images at
https://hub.docker.com/r/lpenz/ - all releases of Debian since Potato
are there, for various architectures. The status of each combination
can be seen [below](#image-status) ([Debian](#debian) and
[Ubuntu](#Ubuntu)).


## Organization

The following files deal with image creation:
- [docker-create-debian-image](docker-create-debian-image): shell
  script that creates a docker image for a specific Debian or Ubuntu
  release, architecture and debootstrap variant.
- [travis-script.sh](travis-script.sh): script that transform a
  travis-ci environment into a call to docker-create-debian-image.


The following deal with distribution-version discovery, updating and rendering:
- [apt-mirror-info](apt-mirror-info): scraps Debian and Ubuntu
  repositories and outputs a json with information about all releases
  it can find.
- [json-tmpl-render](json-tmpl-render): renders a template file with
  information from a json file.
- [README.md.tmpl](README.md.tmpl): template for README.md that uses
  the information obtained by apt-mirror-info to create a table of
  images and status'.
- [SConstruct](SConstruct): scons script that builds the go sources
  and README.md.


Besides image building and deploying to
[docker hub](https://hub.docker.com), the [.travis.yml](.travis.yml)
file also performs static analysis, builds go sources and checks if
the README.md file is up-to-date.


## Image status



### Debian

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>potato</td><td>2.2r7</td><td>alpha</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-alpha.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-alpha">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-alpha.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-alpha-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-alpha-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-alpha-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>arm</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-arm.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-arm">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-arm.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-arm-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-arm-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-arm-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>m68k</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-m68k.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-m68k">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-m68k.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-m68k-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-m68k-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-m68k-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-potato-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>alpha</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-alpha.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-alpha">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-alpha.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-alpha-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-alpha-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-alpha-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>arm</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-arm.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-arm">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-arm.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-arm-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-arm-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-arm-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>m68k</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-m68k.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-m68k">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-m68k.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-m68k-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-m68k-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-m68k-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>s390</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-s390.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-s390">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-s390.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-s390-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-s390-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-s390-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-woody-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>alpha</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-alpha.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-alpha">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-alpha.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-alpha-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-alpha-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-alpha-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>arm</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-arm.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-arm">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-arm.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-arm-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-arm-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-arm-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>m68k</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-m68k.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-m68k">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-m68k.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-m68k-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-m68k-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-m68k-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>s390</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-s390.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-s390">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-s390.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-s390-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-s390-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-s390-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sarge-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>alpha</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-alpha.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-alpha">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-alpha.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-alpha-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-alpha-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-alpha-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>arm</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-arm.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-arm">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-arm.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-arm-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-arm-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-arm-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>s390</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-s390.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-s390">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-s390.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-s390-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-s390-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-s390-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-etch-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>alpha</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-alpha.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-alpha">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-alpha.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-alpha-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-alpha-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-alpha-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>arm</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-arm.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-arm">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-arm.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-arm-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-arm-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-arm-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>s390</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-s390.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-s390">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-s390.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-s390-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-s390-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-s390-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-lenny-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>s390</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-s390.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-s390">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-s390.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-s390-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-s390-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-s390-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-squeeze-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>s390</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-s390.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-s390.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-s390-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-s390-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-wheezy-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-jessie-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-jessie-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mips64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-mips64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-mips64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-stretch-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>mips64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mips64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-mips64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mips64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-mips64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>buster</td><td></td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-buster-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mips64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mips64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-mips64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-mips64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mips64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-mips64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-mips64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>experimental</td><td></td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-experimental-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-experimental-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>mips</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mips.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-mips">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mips.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mips-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-mips-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mips-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>mips64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mips64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-mips64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mips64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mips64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-mips64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mips64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>mipsel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mipsel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-mipsel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mipsel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mipsel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-mipsel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mipsel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>sid</td><td></td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/debian-sid-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-s390x-minbase.svg" />
    </a>
</td>
</tr>
</tbody>
</table>


### Ubuntu

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>warty</td><td>4.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-warty-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-warty-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-warty-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-warty-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-warty-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-warty-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hoary-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-breezy-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>lpia</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-lpia.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-lpia">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-lpia.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-lpia-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-lpia-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-lpia-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>lpia</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-lpia.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-lpia">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-lpia.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-lpia-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-lpia-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-lpia-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>lpia</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-lpia.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-lpia">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-lpia.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-lpia-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-lpia-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-lpia-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>hppa</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-hppa.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-hppa">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-hppa.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-hppa-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-hppa-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-hppa-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>lpia</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-lpia.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-lpia">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-lpia.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-lpia-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-lpia-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-lpia-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>lpia</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-lpia.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-lpia">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-lpia.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-lpia-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-lpia-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-lpia-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>ia64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-ia64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-ia64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-ia64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-ia64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-ia64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-ia64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>sparc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-sparc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-sparc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-sparc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-sparc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-sparc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-sparc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-maverick-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-maverick-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-maverick-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-maverick-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-maverick-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-maverick-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-maverick-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-maverick-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-natty-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-natty-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-natty-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-natty-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-natty-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-natty-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-natty-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-natty-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-oneiric-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-oneiric-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-oneiric-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-oneiric-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-oneiric-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-oneiric-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-oneiric-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-oneiric-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-precise-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-trusty-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-saucy-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-raring-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-raring-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-raring-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-raring-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-raring-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-raring-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-raring-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-raring-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>armel</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-armel.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armel">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armel.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-armel-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armel-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armel-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-quantal-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-utopic-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-vivid-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-wily-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>powerpc</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-powerpc.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-powerpc">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-powerpc.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-powerpc-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-powerpc-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-powerpc-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-s390x-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>amd64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-amd64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-amd64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-amd64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-amd64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-amd64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-amd64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>arm64</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-arm64.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-arm64">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-arm64.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-arm64-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-arm64-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-arm64-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>armhf</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-armhf.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-armhf">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-armhf.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-armhf-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-armhf-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-armhf-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>i386</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-i386.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-i386">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-i386.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-i386-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-i386-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-i386-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>ppc64el</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-ppc64el.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-ppc64el">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-ppc64el.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-ppc64el-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-ppc64el-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-ppc64el-minbase.svg" />
    </a>
</td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>s390x</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-s390x.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-s390x">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-s390x.svg" />
    </a>
</td>
<td>
    <a href="https://travis-ci.com/lpenz/docker-debian-releases">
        <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-s390x-minbase.svg" />
    </a>
    <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-s390x-minbase">
        <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-s390x-minbase.svg" />
    </a>
</td>
</tr>
</tbody>
</table>

