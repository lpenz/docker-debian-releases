[![Build Status](https://img.shields.io/travis/com/lpenz/docker-debian-releases/master.svg?label=master)](https://travis-ci.com/lpenz/docker-debian-releases)


docker-debian-releases
======================

This repository creates docker images of Debian-based system using
debootstrap, for various architectures, and uploads them
to [docker hub](https://hub.docker.com/r/lpenz/) using travis.

The status of each combination is in the tables below:
- [Debian](#Debian)
- [Devuan](#Devuan)
- [Raspbian](#Raspbian)
- [Tanglu](#Tanglu)
- [Ubuntu](#Ubuntu)


## Organization

To avoid having to track the combinations of each distribution and
architecture manually, this repository gets the parameters of
debootstrap from the current branch name, and then scraps a list of
mirrors to figure out which ones to use. That way, to support a new
release, we have to simply push a new remote branch on top of HEAD.
The following scripts are in charge of this mechanism:
- [docker-create-debian-image](docker-create-debian-image): shell
  script that creates a docker image for a specific Debian or Ubuntu
  release, architecture and debootstrap variant.
- [travis-script.sh](travis-script.sh): script that transform a
  travis-ci environment into a call to docker-create-debian-image.
- [dockerhub-set-descriptions](go/cmd/dockerhub-set-descriptions/main.go):
  updates the short and long description of the image in docker hub
  after a new version is deployed.


The README.md file (aka this file) is created offline, from
information obtained from scrapping all available mirrors and
travis-ci itself. The following scripts are in charge of this process:
- [apt-mirror-info](go/cmd/apt-mirror-info/main.go): scraps Debian
  and Ubuntu repositories and outputs a json with information about
  all releases it can find.
- [json-tmpl-render](go/cmd/json-tmpl-render/main.go): renders a
  template file with information from a json file.
- [travis-branch-jobs](go/cmd/travis-branch-jobs/main.go): scraps the
  build information form this repository form travis-ci, for all
  branches that corresponds to relases.
- [README.md.tmpl](README.md.tmpl): template for README.md that uses
  the information obtained by apt-mirror-info and travis-branch-jobs
  to create a table of images and status' with links to jobs.
- [SConstruct](SConstruct): scons script that builds the go sources
  and README.md.


Besides image building and deploying to
[docker hub](https://hub.docker.com), the [.travis.yml](.travis.yml)
file also performs static analysis, builds go sources and checks if
the README.md file is up-to-date.


## Image status

The tables below detail the result of the latest build attempt, and
links to the image in [docker hub](https://hub.docker.com/r/lpenz/) if
the build was successful.

The errors reported below are usually caused by:
- lack of support in qemu for the architecture;
- timeout when building the standard image (that's why minbase is also built);
- incompatibility with modern linux kernel.



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
        <a href="https://hub.docker.com/r/lpenz/debian-potato-alpha">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337577">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-arm">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337585">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337597">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337603">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337608">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337612">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337615">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337619">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337623">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-alpha">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338208">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>arm</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338212">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338214">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338218">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338220">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338231">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338239">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338243">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338247">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338252">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338255">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338258">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338266">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338272">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338279">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338284">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338286">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338287">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338291">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338294">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-alpha">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337631">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-arm">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337640">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337644">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337645">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337655">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337660">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337665">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337671">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337682">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337692">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337707">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337716">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337720">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337723">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337744">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337752">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-alpha">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-alpha-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-alpha-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-arm">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-arm-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-arm-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337098">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337100">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337111">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337118">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mips-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337161">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337167">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337175">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337182">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-alpha">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-alpha-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-alpha-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-arm">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-arm-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-arm-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337464">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337468">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337495">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337498">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mips-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337536">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337545">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337551">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337554">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337903">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337907">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mips-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337958">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337962">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337965">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337975">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338095">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338100">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mips-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338162">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338170">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338190">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338199">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mips-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337995">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113217917">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338015">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338022">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338039">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338043">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113338045">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113336952">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113336959">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113336966">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113336972">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113336975">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113336980">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113336984">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337004">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337019">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337026">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337030">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337041">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337049">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337052">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337057">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337061">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337186">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337192">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337198">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337202">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337212">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337222">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337227">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337235">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337242">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337250">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337259">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337263">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337266">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337268">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337273">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337287">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337291">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337298">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337305">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337316">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337763">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337766">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337773">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337776">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337780">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337781">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337787">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337797">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337802">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337808">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337813">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337821">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337844">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337848">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337850">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337857">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113337863">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
</tbody>
</table>

### Devuan

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>ascii</td><td>2.0</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258665">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258672">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258931">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258939">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>1.0</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-jessie-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-jessie-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-jessie-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-jessie-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>1.0</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259210">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259213">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>1.0</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-jessie-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-jessie-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-jessie-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-jessie-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>1.0</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-jessie-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-jessie-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-jessie-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-jessie-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>1.0</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-jessie-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-jessie-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-jessie-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-jessie-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>1.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259269">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259275">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259005">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259015">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259022">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259027">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259034">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259038">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259047">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259053">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259082">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259086">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259137">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259140">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258563">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113215624">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258583">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258591">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258603">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258613">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258623">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258633">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258643">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258647">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258655">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258659">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258665">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258672">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258682">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258688">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258696">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258703">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258712">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258722">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258799">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258807">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258823">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258830">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>m32</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258845">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258852">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258855">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258865">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258870">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258875">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258897">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258900">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>or1k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258908">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258912">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258917">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258924">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258931">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258939">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258945">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258950">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258955">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258958">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0.0</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258971">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258980">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258987">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113258994">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259005">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259015">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259022">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259027">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259034">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259038">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259047">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259053">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259067">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259074">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259082">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259086">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259088">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259091">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259100">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259107">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259115">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259120">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259126">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259134">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259137">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259140">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259147">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259155">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259164">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113259167">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235047">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235050">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235051">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235054">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235056">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235062">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235064">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235067">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235069">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235071">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235073">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235076">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235077">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235079">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235083">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235084">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235085">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235086">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235091">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235092">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235095">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235099">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235100">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235101">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235105">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235108">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235112">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113235114">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
</tbody>
</table>

### Raspbian

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>jessie</td><td></td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-jessie-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/raspbian-jessie-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-jessie-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/raspbian-jessie-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td></td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-stretch-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/raspbian-stretch-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-stretch-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/raspbian-stretch-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113217952">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-buster-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/raspbian-buster-armhf-minbase.svg" />
        </a>
    </td>
</tr>
</tbody>
</table>

### Tanglu

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>aequorea</td><td>1.0</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-aequorea-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-aequorea-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-aequorea-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-aequorea-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>aequorea</td><td>1.0</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-aequorea-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-aequorea-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-aequorea-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-aequorea-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bartholomea</td><td>2</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-bartholomea-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-bartholomea-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-bartholomea-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-bartholomea-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bartholomea</td><td>2</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-bartholomea-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-bartholomea-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-bartholomea-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-bartholomea-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chromodoris</td><td>3</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-chromodoris-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-chromodoris-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-chromodoris-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-chromodoris-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chromodoris</td><td>3</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-chromodoris-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-chromodoris-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-chromodoris-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-chromodoris-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dasyatis</td><td>3.0</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-dasyatis-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-dasyatis-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113260028">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dasyatis</td><td>3.0</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/tanglu-dasyatis-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/tanglu-dasyatis-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113260048">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>staging</td><td></td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113260053">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113260055">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>staging</td><td></td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113260061">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113260077">
            <img alt="failed" src="build-failed.svg" />
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382552">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382560">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382537">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382543">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382564">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382570">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380942">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380946">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380930">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380938">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380975">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380977">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380953">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380956">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380991">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380998">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379948">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379954">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380035">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379928">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379979">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379987">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379962">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379970">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379995">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380007">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379937">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379941">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380149">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380159">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380236">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380249">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380224">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380230">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380187">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380194">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380204">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380217">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380166">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380173">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380347">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380362">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380370">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380379">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380388">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380399">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380406">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380412">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380417">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380424">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380437">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380447">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380583">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380588">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380602">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380613">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380620">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380631">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380646">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380652">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380659">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380712">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380695">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380732">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380739">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-lpia">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-lpia-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380786">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380797">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380833">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380834">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380853">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380867">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-lpia">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-lpia-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380916">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380920">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381018">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381028">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381064">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381076">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-lpia">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-lpia-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381130">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381149">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381195">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381197">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381237">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381246">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-lpia">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-lpia-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381298">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381313">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381392">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381403">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-lpia">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-lpia-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381453">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381464">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381549">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381562">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381584">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113381589">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382768">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382778">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382787">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382792">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-powerpc-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382908">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382918">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382929">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113382940">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113383005">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113383008">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113383016">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113383020">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113218003">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113218005">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113218007">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113218009">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379877">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379892">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379899">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113379924">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380061">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380068">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380111">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380116">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380128">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380137">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380318">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380325">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380336">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380341">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380537">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380542">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380549">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380554">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380537">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380542">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380549">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/113380554">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>

</tbody>
</table>
