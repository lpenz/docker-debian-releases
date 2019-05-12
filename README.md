[![Build Status](https://img.shields.io/travis/com/lpenz/docker-debian-releases/master.svg?label=master)](https://travis-ci.com/lpenz/docker-debian-releases)


docker-debian-releases
======================

This repository creates docker images of Debian-based system using
debootstrap, for various architectures, and uploads them
to [docker hub](https://hub.docker.com/r/lpenz/) using travis.

The status of each combination is in the tables below:
- [Debian](#debian)
- [Raspbian](#raspbian)
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
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766117">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-alpha-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-arm">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766127">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-arm-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766137">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785481">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-m68k.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766140">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-m68k-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785492">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766148">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785501">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766159">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-potato-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-alpha">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766167">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-alpha-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>arm</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785913">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766177">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-arm-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785926">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766183">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766193">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785952">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766207">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785955">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-m68k.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766222">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-m68k-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785964">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766228">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785970">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766235">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785976">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766262">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785985">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-s390.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766253">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-s390-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785991">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766269">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-woody-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-alpha">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766276">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-alpha-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-arm">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766282">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-arm-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785520">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766289">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766307">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785539">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766318">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785556">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-m68k.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766333">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-m68k-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766342">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766346">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766350">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785590">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-s390.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766358">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-s390-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785602">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766367">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sarge-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-alpha">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-alpha-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-alpha-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-arm">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-arm-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-arm-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785185">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110844990">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785217">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845021">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mips-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mipsel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-etch-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785244">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-s390.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845050">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-s390-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785432">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845059">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-etch-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-alpha">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-alpha.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-alpha-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-alpha-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-arm">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-arm-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-arm-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785404">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845188">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785421">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845209">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mips-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mipsel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-lenny-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785449">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-s390.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845238">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-s390-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785457">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845256">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-lenny-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785649">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845294">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mips-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mipsel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-squeeze-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785683">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-s390.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845342">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-s390-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785690">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845350">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-squeeze-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785830">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845477">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mips-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mipsel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785883">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-s390.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845532">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-s390-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390x">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390x-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-wheezy-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785902">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845549">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-wheezy-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mips-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mipsel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-ppc64el">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-ppc64el-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-s390x">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-s390x-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-jessie-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785724">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785735">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785763">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-mips64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845407">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-mips64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mipsel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785781">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845420">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.9</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110794986">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-stretch-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-s390x-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766378">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770781">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766382">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770795">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766385">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770804">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766395">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766411">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766414">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770838">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mips64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766420">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mips64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766428">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770860">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766438">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770864">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766456">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-buster-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770870">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766461">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770880">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766467">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770894">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766477">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770898">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766486">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770907">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766492">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770911">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766497">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770918">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mips64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766507">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mips64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770924">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766513">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770928">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766531">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770933">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766541">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-experimental-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766545">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770959">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766551">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770966">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766556">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770968">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766563">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766572">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-mips">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766594">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110770987">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mips64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766598">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mips64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-mipsel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766612">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110771010">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766626">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110771019">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110766637">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/debian-sid-s390x-minbase.svg" />
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
    <td>wheezy</td><td>7.0</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-wheezy-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/raspbian-wheezy-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-wheezy-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/raspbian-wheezy-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td></td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-jessie-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/raspbian-jessie-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-jessie-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/raspbian-jessie-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td></td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-stretch-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/raspbian-stretch-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-stretch-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/raspbian-stretch-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110945644">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/raspbian-buster-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-buster-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/raspbian-buster-armhf-minbase.svg" />
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
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845715">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795220">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845708">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795237">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845719">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-warty-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795189">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845678">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795176">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845662">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795206">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845692">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795197">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845686">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795213">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845695">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hoary-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795022">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785026">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795001">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785014">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795044">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845569">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795032">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785032">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795053">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845578">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795014">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110785024">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-breezy-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795064">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845586">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795114">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845614">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795094">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845611">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795076">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845600">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795083">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845607">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795072">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845594">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-dapper-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795123">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845618">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795131">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845626">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795139">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845631">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795148">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845641">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795154">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845650">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110795162">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110845652">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-edgy-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739005">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110968042">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739014">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739035">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110968376">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739036">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739042">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110968733">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739054">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-feisty-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110969093">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739073">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110969465">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739089">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-lpia">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-lpia-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-gutsy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110969996">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739124">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-gutsy-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110970360">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739148">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110970746">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739171">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-lpia">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-lpia-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-hardy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110971285">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739207">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-hardy-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110971571">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739225">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110971935">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739241">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-lpia">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-lpia-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-intrepid-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110972445">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739280">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-intrepid-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110973002">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-hppa.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739303">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-hppa-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110973317">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739325">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-lpia">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-lpia-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-jaunty-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110973765">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739351">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-jaunty-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110974327">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739390">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-lpia">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-lpia.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-lpia-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-lpia-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-karmic-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110974711">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739437">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-karmic-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110975356">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-ia64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739486">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-ia64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-lucid-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110975585">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-sparc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110739508">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-lucid-sparc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-maverick-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-natty-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-oneiric-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-precise-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-ppc64el">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-ppc64el-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-trusty-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-saucy-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-raring-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armel">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armel-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-quantal-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-ppc64el">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-ppc64el-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-utopic-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-ppc64el">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-ppc64el-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-vivid-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-ppc64el">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-ppc64el-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-wily-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110982454">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740230">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110982573">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740241">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-xenial-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-powerpc">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-powerpc-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-powerpc-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110983295">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740320">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110983395">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740329">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-yakkety-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110983948">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740389">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110984060">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740390">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-zesty-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-artful-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950848">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740431">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950853">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740437">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-artful-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950891">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740485">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950906">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740506">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-bionic-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950940">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740564">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950957">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740584">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110950961">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740598">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-cosmic-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-amd64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-amd64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-arm64">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-arm64-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-armhf">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-armhf-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-i386">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-i386-minbase">
            <img alt="" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110951020">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740635">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110951025">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740651">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-disco-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110966673">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740659">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110966882">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740681">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967073">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740691">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967283">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740699">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967508">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740716">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110967688">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/110740725">
            <img alt="" src="https://img.shields.io/travis/com/lpenz/docker-debian-releases/ubuntu-eoan-s390x-minbase.svg" />
        </a>
    </td>
</tr>
</tbody>
</table>

