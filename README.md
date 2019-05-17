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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766117">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766127">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766137">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785481">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766140">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785492">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766148">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785501">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766159">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766167">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>arm</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785913">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766177">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785926">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766183">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766193">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785952">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766207">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785955">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766222">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785964">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766228">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785970">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766235">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785976">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766262">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785985">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766253">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785991">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766269">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766276">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766282">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785520">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766289">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766307">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785539">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766318">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785556">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766333">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766342">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766346">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766350">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785590">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766358">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785602">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766367">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785185">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110844990">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785217">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845021">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785244">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845050">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785432">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845059">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785404">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845188">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785421">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845209">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785449">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845238">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785457">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845256">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785649">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845294">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785683">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845342">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785690">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845350">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785830">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845477">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785883">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845532">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785902">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845549">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785724">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785735">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785763">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845407">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785781">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845420">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110794986">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766378">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770781">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766382">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770795">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766385">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770804">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766395">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766411">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766414">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770838">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766420">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766428">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770860">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766438">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770864">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766456">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770870">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766461">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770880">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766467">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770894">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766477">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770898">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766486">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770907">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766492">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770911">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766497">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770918">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766507">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770924">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766513">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770928">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766531">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770933">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766541">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766545">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770959">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766551">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770966">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766556">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770968">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766563">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766572">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766594">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770987">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766598">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766612">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110771010">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766626">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110771019">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766637">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984700">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984713">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984773">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984781">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111983996">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984003">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984011">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984027">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984029">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984033">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984037">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984042">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984071">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984079">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984140">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984147">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984202">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984208">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984223">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984231">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984234">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984236">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984245">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984249">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984282">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984289">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984411">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984421">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984504">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984509">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984513">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984521">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984528">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984536">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984554">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984561">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984578">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984584">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984633">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/111984641">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110945644">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795231">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845715">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795220">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845708">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795237">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845719">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795189">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845678">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795176">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845662">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795206">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845692">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795197">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845686">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795213">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845695">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795022">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785026">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795001">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785014">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795044">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845569">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795032">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785032">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795053">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845578">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795014">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785024">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795064">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845586">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795114">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845614">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795094">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845611">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795076">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845600">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795083">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845607">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795072">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845594">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795123">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845618">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795131">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845626">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795139">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845631">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795148">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845641">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795154">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845650">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795162">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845652">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739005">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110968042">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739014">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739035">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110968376">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739036">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739042">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110968733">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739054">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110969093">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739073">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110969465">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739089">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110969996">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739124">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110970360">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739148">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110970746">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739171">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110971285">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739207">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110971571">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739225">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110971935">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739241">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110972445">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739280">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110973002">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739303">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110973317">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739325">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110973765">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739351">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110974327">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739390">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110974711">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739437">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110975356">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739486">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110975585">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739508">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110982454">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740230">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110982573">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740241">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110983295">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740320">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110983395">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740329">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110983948">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740389">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110984060">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740390">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950848">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740431">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950853">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740437">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950891">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740485">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950906">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740506">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950940">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740564">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950957">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740584">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950961">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740598">
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110951020">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740635">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110951025">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740651">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110966673">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740659">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110966882">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740681">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967073">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740691">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967283">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740699">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967508">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740716">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967688">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740725">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110966673">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740659">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110966882">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740681">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967073">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740691">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967283">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740699">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967508">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740716">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967688">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740725">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>

</tbody>
</table>
