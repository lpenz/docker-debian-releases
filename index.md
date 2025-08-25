docker-debian-releases
======================

The
[docker-debian-releases](https://github.com/lpenz/docker-debian-releases)
repository creates docker images of Debian-based system using
debootstrap, for various architectures, and uploads them to [docker
hub](https://hub.docker.com/r/lpenz/) using github actions.

The tables below detail the result of the latest build attempt, and
links to the image in [docker hub](https://hub.docker.com/r/lpenz/) if
the build was successful. The supported OS's are:

- [Debian](#debian)
- [Devuan](#devuan)
- [Raspbian](#raspbian)
- [Ubuntu](#ubuntu)
- [rpios](#rpios)


Most of the errors are caused by:
- lack of support in qemu for the architecture;
- timeout when building the standard image (that's why minbase is also built);
- incompatibility with modern linux kernel.



### Debian

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Release date</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-alpha">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-alpha">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-alpha-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-alpha-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-arm">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-arm">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-arm-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-arm-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>m68k</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-m68k">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-m68k">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-m68k-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-m68k-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-potato-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-alpha">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-alpha">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-alpha-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-alpha-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-arm">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-arm">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-arm-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-arm-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>m68k</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-m68k">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-m68k">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-m68k-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-m68k-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-mips">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-mips-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>s390</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-s390">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-s390">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-s390-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-s390-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-woody-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-alpha">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-alpha">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-alpha-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-alpha-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-arm">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-arm">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-arm-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-arm-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>m68k</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-m68k">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-m68k">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-m68k-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-m68k-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-mips">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-mips-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>s390</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-s390">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-s390">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-s390-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-s390-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sarge-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-alpha">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-alpha">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-alpha-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-alpha-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-arm">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-arm">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-arm-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-arm-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-mips">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-mips-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>s390</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-s390">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-s390">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-s390-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-s390-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-etch-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-etch-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>alpha</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-alpha">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-alpha">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-alpha-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-alpha-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-arm">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-arm">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-arm-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-arm-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-mips">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-mips-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>s390</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-s390">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-s390">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-s390-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-s390-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-lenny-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-lenny-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-mips">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-mips-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>s390</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-s390">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-s390">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-s390-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-s390-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-squeeze-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-squeeze-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-armel">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-armel-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-armhf">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-armhf-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-i386">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-i386-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-ia64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-ia64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-mips">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-mips-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-mipsel">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-mipsel-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-powerpc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-powerpc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>s390</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-s390">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-s390-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-s390x">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-s390x-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-sparc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-wheezy-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-wheezy-sparc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-mips">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-mips-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-jessie-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-jessie-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-mips">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-mips-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>mips64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-mips64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-mips64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2021-08-14T07:42:00Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-stretch-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-mips">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-mips-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>mips64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-mips64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-mips64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.13</td><td>2023-06-10T08:53:33Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-buster-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>mips64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-mips64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-mips64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-mips64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-mips64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2024-08-31T11:02:15Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>mips64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-mips64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-mips64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-mips64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-mips64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td>12.11</td><td>2025-08-09T08:53:07Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bookworm-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bookworm-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2025-08-09T09:16:30Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2025-08-09T09:16:30Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2025-08-09T09:16:30Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td>11.11</td><td>2025-08-09T09:16:30Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-bullseye-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td>13.0</td><td>2025-08-09T10:59:34Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td>13.0</td><td>2025-08-09T10:59:34Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td>13.0</td><td>2025-08-09T10:59:34Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td>13.0</td><td>2025-08-09T10:59:34Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td>13.0</td><td>2025-08-09T10:59:34Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td>13.0</td><td>2025-08-09T10:59:34Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td>13.0</td><td>2025-08-09T10:59:34Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td>13.0</td><td>2025-08-09T10:59:34Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-trixie-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-trixie-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>forky</td><td></td><td>2025-08-24T14:28:48Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>forky</td><td></td><td>2025-08-24T14:28:48Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-arm64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-arm64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>forky</td><td></td><td>2025-08-24T14:28:48Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-armel">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-armel-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>forky</td><td></td><td>2025-08-24T14:28:48Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-armhf">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-armhf-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>forky</td><td></td><td>2025-08-24T14:28:48Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-i386">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-i386-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>forky</td><td></td><td>2025-08-24T14:28:48Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-ppc64el">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-ppc64el-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>forky</td><td></td><td>2025-08-24T14:28:48Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-riscv64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-riscv64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>forky</td><td></td><td>2025-08-24T14:28:48Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-s390x">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-forky-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-forky-s390x-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>mips64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-mips64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-mips64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-mips64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-mips64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>rc-buggy</td><td></td><td>2025-08-24T14:28:48Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-rc-buggy-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-rc-buggy-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>mips64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-mips64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-mips64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-mips64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-mips64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2025-08-24T14:28:48Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adebian-sid-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
</tbody>
</table>

### Devuan

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Release date</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>chimaera</td><td>4.0</td><td>2025-08-24T03:22:50Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0</td><td>2025-08-24T03:22:50Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0</td><td>2025-08-24T03:22:50Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0</td><td>2025-08-24T03:22:50Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0</td><td>2025-08-24T03:22:50Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0</td><td>2025-08-24T03:22:50Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-24T03:21:55Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-24T03:21:55Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-24T03:21:55Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-24T03:21:55Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-24T03:21:55Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-24T03:21:55Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-24T03:23:53Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-24T03:23:53Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-24T03:23:53Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-24T03:23:53Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-24T03:23:53Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-24T03:23:53Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-riscv64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-riscv64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-24T03:23:53Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-24T15:04:11Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-24T15:04:11Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-24T15:04:11Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-24T15:04:11Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-24T15:04:11Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-24T15:04:11Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2025-08-25T00:15:06Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2025-08-25T00:15:06Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2025-08-25T00:15:06Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2025-08-25T00:15:06Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2025-08-25T00:15:06Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2025-08-25T00:15:06Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-mipsel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-mipsel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-mipsel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-mipsel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2025-08-25T00:15:06Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ascii-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2025-08-25T00:15:05Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2025-08-25T00:15:05Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2025-08-25T00:15:05Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2025-08-25T00:15:05Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2025-08-25T00:15:05Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2025-08-25T00:15:05Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-beowulf-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2025-08-25T00:15:05Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2025-08-25T00:15:05Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2025-08-25T00:15:05Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2025-08-25T00:15:05Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2025-08-25T00:15:05Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2025-08-25T00:15:05Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-chimaera-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-chimaera-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-25T00:15:06Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-25T00:15:06Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-25T00:15:06Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-25T00:15:06Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-25T00:15:06Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>daedalus</td><td>5.0</td><td>2025-08-25T00:15:06Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-daedalus-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-daedalus-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-25T00:15:06Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-25T00:15:06Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-25T00:15:06Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-25T00:15:06Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-25T00:15:06Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-25T00:15:06Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>excalibur</td><td>6.0</td><td>2025-08-25T00:15:06Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-riscv64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-excalibur-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-excalibur-riscv64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>freia</td><td>7.0</td><td>2025-08-25T00:15:06Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-amd64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-amd64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>freia</td><td>7.0</td><td>2025-08-25T00:15:06Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-arm64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-arm64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>freia</td><td>7.0</td><td>2025-08-25T00:15:06Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-armel">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-armel-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>freia</td><td>7.0</td><td>2025-08-25T00:15:06Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-armhf">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-armhf-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>freia</td><td>7.0</td><td>2025-08-25T00:15:06Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-i386">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-i386-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>freia</td><td>7.0</td><td>2025-08-25T00:15:06Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-ppc64el">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-ppc64el-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>freia</td><td>7.0</td><td>2025-08-25T00:15:06Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-riscv64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-freia-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-freia-riscv64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-25T00:15:06Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-25T00:15:06Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-25T00:15:06Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-25T00:15:06Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-25T00:15:06Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-25T00:15:06Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2025-08-25T00:15:06Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-ceres-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2025-08-25T00:15:05Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2025-08-25T00:15:05Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2025-08-25T00:15:05Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2025-08-25T00:15:05Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2025-08-25T00:15:05Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2025-08-25T00:15:05Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2025-08-25T00:15:05Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-experimental-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Adevuan-experimental-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
</tbody>
</table>

### Raspbian

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Release date</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>buster</td><td></td><td>2025-07-13T04:16:03Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-buster-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Araspbian-buster-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-buster-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Araspbian-buster-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2025-08-24T22:16:49Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-bullseye-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Araspbian-bullseye-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-bullseye-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Araspbian-bullseye-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td></td><td>2025-08-24T22:23:12Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-bookworm-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Araspbian-bookworm-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-bookworm-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Araspbian-bookworm-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td></td><td>2025-08-24T22:31:11Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-trixie-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Araspbian-trixie-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-trixie-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Araspbian-trixie-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
</tbody>
</table>

### Ubuntu

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Release date</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>warty</td><td>4.10</td><td>2005-05-13T12:42:31Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-warty-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-warty-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>2005-05-13T12:42:31Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-warty-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-warty-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>2005-05-13T12:42:31Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-warty-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-warty-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-warty-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hoary-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hoary-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-breezy-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-breezy-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-dapper-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-edgy-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-feisty-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-lpia">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-lpia">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-lpia-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-lpia-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-gutsy-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-gutsy-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-hppa">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-hppa-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-i386">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-i386-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-ia64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-ia64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-lpia">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-lpia">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-lpia-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-lpia-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-powerpc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-powerpc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-sparc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hardy-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hardy-sparc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-hppa">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-hppa-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-i386">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-i386-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-ia64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-ia64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-lpia">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-lpia">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-lpia-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-lpia-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-powerpc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-powerpc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-sparc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-intrepid-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-intrepid-sparc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>hppa</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-hppa">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-hppa">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-hppa-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-hppa-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-lpia">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-lpia">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-lpia-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-lpia-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jaunty-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jaunty-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-ia64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-ia64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>lpia</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-lpia">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-lpia">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-lpia-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-lpia-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-sparc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-karmic-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-karmic-sparc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-armel">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-armel-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-i386">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-i386-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>ia64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-ia64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-ia64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-ia64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-ia64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-powerpc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-powerpc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>sparc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-sparc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-sparc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lucid-sparc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lucid-sparc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>2010-10-10T10:18:49Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-maverick-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-maverick-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>2010-10-10T10:18:49Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-maverick-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-maverick-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>2010-10-10T10:18:49Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-maverick-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-maverick-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>2010-10-10T10:18:49Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-maverick-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-maverick-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-maverick-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>2011-04-26T12:16:58Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-natty-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-natty-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>2011-04-26T12:16:58Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-natty-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-natty-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>2011-04-26T12:16:58Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-natty-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-natty-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>natty</td><td>11.04</td><td>2011-04-26T12:16:58Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-natty-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-natty-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-natty-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>2011-10-12T18:19:26Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oneiric-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oneiric-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>2011-10-12T18:19:26Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oneiric-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oneiric-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>2011-10-12T18:19:26Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oneiric-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oneiric-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oneiric</td><td>11.10</td><td>2011-10-12T18:19:26Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oneiric-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oneiric-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oneiric-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-armel">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-armel-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-armhf">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-armhf-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-i386">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-i386-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-powerpc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-precise-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-precise-powerpc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-trusty-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-trusty-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-saucy-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-saucy-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>2014-05-08T14:19:55Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-raring-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-raring-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>2014-05-08T14:19:55Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-raring-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-raring-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>2014-05-08T14:19:55Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-raring-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-raring-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>raring</td><td>13.04</td><td>2014-05-08T14:19:55Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-raring-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-raring-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-raring-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-armel">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-armel-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-quantal-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-quantal-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-utopic-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-utopic-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-arm64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-arm64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-armhf">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-armhf-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-i386">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-i386-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-powerpc">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-powerpc-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-ppc64el">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-vivid-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-vivid-ppc64el-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-wily-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-wily-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-xenial-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-powerpc">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-powerpc">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-powerpc-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-powerpc-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-yakkety-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-zesty-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-artful-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-artful-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-bionic-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-cosmic-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-disco-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-arm64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-arm64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-armhf">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-armhf-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-i386">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-i386-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-ppc64el">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-ppc64el-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-s390x">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-eoan-s390x-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-focal-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-groovy-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-04-22T16:33:50Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-04-22T16:33:50Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-04-22T16:33:50Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-04-22T16:33:50Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-04-22T16:33:50Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-04-22T16:33:50Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-04-22T16:33:50Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-hirsute-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-hirsute-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>impish</td><td>21.10</td><td>2021-10-14T16:35:09Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>impish</td><td>21.10</td><td>2021-10-14T16:35:09Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-arm64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-arm64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>impish</td><td>21.10</td><td>2021-10-14T16:35:09Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-armhf">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-armhf-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>impish</td><td>21.10</td><td>2021-10-14T16:35:09Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-i386">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-i386-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>impish</td><td>21.10</td><td>2021-10-14T16:35:09Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-ppc64el">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-ppc64el-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>impish</td><td>21.10</td><td>2021-10-14T16:35:09Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-riscv64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-riscv64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>impish</td><td>21.10</td><td>2021-10-14T16:35:09Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-s390x">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-impish-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-impish-s390x-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jammy</td><td>22.04</td><td>2022-04-21T17:16:08Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jammy</td><td>22.04</td><td>2022-04-21T17:16:08Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jammy</td><td>22.04</td><td>2022-04-21T17:16:08Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jammy</td><td>22.04</td><td>2022-04-21T17:16:08Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jammy</td><td>22.04</td><td>2022-04-21T17:16:08Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jammy</td><td>22.04</td><td>2022-04-21T17:16:08Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jammy</td><td>22.04</td><td>2022-04-21T17:16:08Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-jammy-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-jammy-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>kinetic</td><td>22.10</td><td>2022-10-20T19:47:56Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-amd64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-amd64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>kinetic</td><td>22.10</td><td>2022-10-20T19:47:56Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-arm64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-arm64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>kinetic</td><td>22.10</td><td>2022-10-20T19:47:56Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-armhf">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-armhf-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>kinetic</td><td>22.10</td><td>2022-10-20T19:47:56Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-i386">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-i386-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>kinetic</td><td>22.10</td><td>2022-10-20T19:47:56Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-ppc64el">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-ppc64el-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>kinetic</td><td>22.10</td><td>2022-10-20T19:47:56Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-riscv64">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-riscv64-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>kinetic</td><td>22.10</td><td>2022-10-20T19:47:56Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-s390x">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-kinetic-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-kinetic-s390x-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lunar</td><td>23.04</td><td>2023-04-20T14:50:49Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lunar</td><td>23.04</td><td>2023-04-20T14:50:49Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lunar</td><td>23.04</td><td>2023-04-20T14:50:49Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lunar</td><td>23.04</td><td>2023-04-20T14:50:49Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lunar</td><td>23.04</td><td>2023-04-20T14:50:49Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lunar</td><td>23.04</td><td>2023-04-20T14:50:49Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lunar</td><td>23.04</td><td>2023-04-20T14:50:49Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-lunar-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-lunar-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>mantic</td><td>23.10</td><td>2023-10-16T21:12:12Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>mantic</td><td>23.10</td><td>2023-10-16T21:12:12Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>mantic</td><td>23.10</td><td>2023-10-16T21:12:12Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>mantic</td><td>23.10</td><td>2023-10-16T21:12:12Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>mantic</td><td>23.10</td><td>2023-10-16T21:12:12Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>mantic</td><td>23.10</td><td>2023-10-16T21:12:12Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>mantic</td><td>23.10</td><td>2023-10-16T21:12:12Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-mantic-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-mantic-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>noble</td><td>24.04</td><td>2024-04-25T15:10:33Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>noble</td><td>24.04</td><td>2024-04-25T15:10:33Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>noble</td><td>24.04</td><td>2024-04-25T15:10:33Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>noble</td><td>24.04</td><td>2024-04-25T15:10:33Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>noble</td><td>24.04</td><td>2024-04-25T15:10:33Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>noble</td><td>24.04</td><td>2024-04-25T15:10:33Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>noble</td><td>24.04</td><td>2024-04-25T15:10:33Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-noble-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-noble-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oracular</td><td>24.10</td><td>2024-10-11T10:20:59Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oracular</td><td>24.10</td><td>2024-10-11T10:20:59Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oracular</td><td>24.10</td><td>2024-10-11T10:20:59Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oracular</td><td>24.10</td><td>2024-10-11T10:20:59Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oracular</td><td>24.10</td><td>2024-10-11T10:20:59Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oracular</td><td>24.10</td><td>2024-10-11T10:20:59Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>oracular</td><td>24.10</td><td>2024-10-11T10:20:59Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-oracular-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-oracular-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>plucky</td><td>25.04</td><td>2025-04-17T15:43:00Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>plucky</td><td>25.04</td><td>2025-04-17T15:43:00Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>plucky</td><td>25.04</td><td>2025-04-17T15:43:00Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>plucky</td><td>25.04</td><td>2025-04-17T15:43:00Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>plucky</td><td>25.04</td><td>2025-04-17T15:43:00Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-ppc64el">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-ppc64el-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>plucky</td><td>25.04</td><td>2025-04-17T15:43:00Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-riscv64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-riscv64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>plucky</td><td>25.04</td><td>2025-04-17T15:43:00Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-s390x">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-plucky-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-plucky-s390x-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>questing</td><td>25.10</td><td>2025-08-25T01:25:29Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-amd64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-amd64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>questing</td><td>25.10</td><td>2025-08-25T01:25:29Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-arm64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-arm64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>questing</td><td>25.10</td><td>2025-08-25T01:25:29Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-armhf">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-armhf-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>questing</td><td>25.10</td><td>2025-08-25T01:25:29Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-i386">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-i386-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>questing</td><td>25.10</td><td>2025-08-25T01:25:29Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-ppc64el">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-ppc64el">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-ppc64el-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-ppc64el-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>questing</td><td>25.10</td><td>2025-08-25T01:25:29Z</td><td>riscv64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-riscv64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-riscv64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-riscv64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-riscv64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>questing</td><td>25.10</td><td>2025-08-25T01:25:29Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-s390x">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-s390x">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-questing-s390x-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Aubuntu-questing-s390x-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
</tbody>
</table>

### rpios

<table>
<thead>
<tr><th rowspan="2">Release</th><th rowspan="2">Version</th><th rowspan="2">Release date</th><th rowspan="2">Arch</th><th colspan="2">Variant status</th></tr>
<tr><th>standard</th><th>minbase</th></tr>
</thead>
<tbody>
<tr>
    <td>wheezy</td><td></td><td>2025-06-05T20:47:08Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-wheezy-armel">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-wheezy-armel">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-wheezy-armel-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-wheezy-armel-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td></td><td>2025-06-05T20:47:08Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-wheezy-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-wheezy-armhf">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-wheezy-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-wheezy-armhf-minbase">
            <img alt="success" src="build-success.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td></td><td>2025-06-05T20:47:09Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-jessie-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-jessie-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-jessie-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-jessie-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td></td><td>2025-06-05T20:47:09Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-jessie-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-jessie-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-jessie-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-jessie-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td></td><td>2025-06-05T20:47:11Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-stretch-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-stretch-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-stretch-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-stretch-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td></td><td>2025-06-05T20:47:11Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-stretch-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-stretch-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-stretch-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-stretch-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td></td><td>2025-06-05T20:47:11Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-stretch-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-stretch-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-stretch-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-stretch-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>2025-08-22T21:31:14Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-buster-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-buster-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-buster-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-buster-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>2025-08-22T21:31:14Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-buster-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-buster-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-buster-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-buster-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>2025-08-22T21:31:14Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-buster-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-buster-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-buster-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-buster-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td></td><td>2025-08-22T21:31:14Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-buster-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-buster-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-buster-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-buster-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2025-08-22T21:31:21Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bullseye-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bullseye-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bullseye-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bullseye-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2025-08-22T21:31:21Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bullseye-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bullseye-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bullseye-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bullseye-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2025-08-22T21:31:21Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bullseye-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bullseye-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bullseye-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bullseye-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2025-08-22T21:31:21Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bullseye-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bullseye-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bullseye-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bullseye-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td></td><td>2025-08-22T21:31:35Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bookworm-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bookworm-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bookworm-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bookworm-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td></td><td>2025-08-22T21:31:35Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bookworm-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bookworm-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bookworm-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bookworm-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td></td><td>2025-08-22T21:31:35Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bookworm-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bookworm-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bookworm-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bookworm-i386-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bookworm</td><td></td><td>2025-08-22T21:31:35Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bookworm-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bookworm-amd64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-bookworm-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-bookworm-amd64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td></td><td>2025-08-22T21:31:41Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-trixie-armhf">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-trixie-armhf">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-trixie-armhf-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-trixie-armhf-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td></td><td>2025-08-22T21:31:41Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-trixie-arm64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-trixie-arm64">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-trixie-arm64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-trixie-arm64-minbase">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td></td><td>2025-08-22T21:31:41Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-trixie-i386">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-trixie-i386">
            <img alt="notfound" src="build-notfound.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-trixie-i386-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-trixie-i386-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>trixie</td><td></td><td>2025-08-22T21:31:41Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-trixie-amd64">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-trixie-amd64">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/rpios-trixie-amd64-minbase">
            <img alt="dockerhub" src="docker-hub.svg" />
        </a>
        <a href="https://github.com/lpenz/docker-debian-releases/actions?query=branch%3Arpios-trixie-amd64-minbase">
            <img alt="failure" src="build-failure.svg" />
        </a>
    </td>
</tr>

</tbody>
</table>
