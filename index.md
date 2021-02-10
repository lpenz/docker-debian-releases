docker-debian-releases
======================

The
[docker-debian-releases](https://github.com/lpenz/docker-debian-releases)
repository creates docker images of Debian-based system using
debootstrap, for various architectures, and uploads them to [docker
hub](https://hub.docker.com/r/lpenz/) using travis.

The tables below detail the result of the latest build attempt, and
links to the image in [docker hub](https://hub.docker.com/r/lpenz/) if
the build was successful. The supported OS's are:

- [Debian](#debian)
- [Devuan](#devuan)
- [Raspbian](#raspbian)
- [Ubuntu](#ubuntu)


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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314566">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314576">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-arm">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314588">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-potato-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-potato-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314596">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314600">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314608">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314613">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314618">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>potato</td><td>2.2r7</td><td>2002-07-12T16:16:28Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314624">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188314626">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937714">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937715">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>arm</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937718">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937719">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937722">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937726">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-woody-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-woody-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937730">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937733">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937735">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937738">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937741">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937744">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937755">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937759">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937760">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937761">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937765">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937766">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937767">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>woody</td><td>3.0r6</td><td>2005-05-31T20:55:05Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937769">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187937771">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321848">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321852">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>arm</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-arm">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-arm.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321867">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321875">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321884">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321900">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321913">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321915">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321919">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321922">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321932">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321947">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sarge-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sarge-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321956">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321962">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321969">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sarge</td><td>3.1r8</td><td>2008-04-12T19:04:58Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321972">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188321981">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939402">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939405">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>amd64</td>
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
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>arm</td>
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
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939420">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939425">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>i386</td>
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
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939433">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939442">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>mips</td>
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
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>mipsel</td>
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
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>powerpc</td>
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
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939462">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939466">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>etch</td><td>4.0r9</td><td>2010-05-22T14:22:09Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939469">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187939471">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994544">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994551">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>amd64</td>
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
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>arm</td>
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
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>armel</td>
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
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994569">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994571">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>i386</td>
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
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994579">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994584">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>mips</td>
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
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>mipsel</td>
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
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>powerpc</td>
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
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994604">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994606">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lenny</td><td>5.0.10</td><td>2012-03-10T11:30:56Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994607">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187994610">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>amd64</td>
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
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>armel</td>
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
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>i386</td>
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
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188510107">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188510119">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>mips</td>
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
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>mipsel</td>
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
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>powerpc</td>
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
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188510173">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188510178">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>squeeze</td><td>6.0.10</td><td>2015-04-25T11:01:14Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188510184">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188510187">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>amd64</td>
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
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>armel</td>
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
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>armhf</td>
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
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>i386</td>
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
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188498819">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188498824">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>mips</td>
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
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>mipsel</td>
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
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>powerpc</td>
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
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188498865">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188498867">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>s390x</td>
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
    <td>wheezy</td><td>7.11</td><td>2017-06-17T08:55:32Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188498890">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188498893">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>amd64</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>arm64</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>armel</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>armhf</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>i386</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>mips</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>mipsel</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>powerpc</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>ppc64el</td>
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
    <td>jessie</td><td>8.11</td><td>2018-06-23T10:30:18Z</td><td>s390x</td>
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
    <td>jessie</td><td>8.11</td><td>2019-07-06T09:36:16Z</td><td>amd64</td>
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
    <td>jessie</td><td>8.11</td><td>2019-07-06T09:36:16Z</td><td>armel</td>
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
    <td>jessie</td><td>8.11</td><td>2019-07-06T09:36:16Z</td><td>armhf</td>
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
    <td>jessie</td><td>8.11</td><td>2019-07-06T09:36:16Z</td><td>i386</td>
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
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>amd64</td>
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
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>arm64</td>
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
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>i386</td>
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
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>mips</td>
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
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>mips64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-mips64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-mips64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>mipsel</td>
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
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>stretch</td><td>9.13</td><td>2020-07-18T10:50:51Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-stretch-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-stretch-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>mips</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>mips64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mips64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mips64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>mipsel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mipsel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mipsel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>buster</td><td>10.8</td><td>2021-02-06T10:10:28Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-buster-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-buster-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>all</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/206926339">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/206926345">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935804">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-bullseye-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935812">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-bullseye-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935819">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-bullseye-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935825">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-bullseye-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935834">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-bullseye-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935842">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-mips64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-bullseye-mips64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935850">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-bullseye-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935859">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-bullseye-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-bullseye-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T20:10:54Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935869">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935878">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>all</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/193059216">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/193059223">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997351">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997354">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997358">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997360">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997361">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997362">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997366">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997368">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997370">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997373">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997375">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997380">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997383">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997384">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997386">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997388">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td></td><td>2021-02-10T20:10:54Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997394">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997399">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>all</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/206929716">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/206929726">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/216385275">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188318108">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188318118">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188318131">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/216385277">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>mips64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188318149">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-mips64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mips64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188318163">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-mipsel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-mipsel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188318177">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/debian-sid-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/debian-sid-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>sid</td><td></td><td>2021-02-10T20:10:54Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188318186">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188318191">
            <img alt="failed" src="build-failed.svg" />
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
    <td>jessie</td><td>1.0</td><td>2021-02-10T03:30:01Z</td><td>amd64</td>
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
    <td>jessie</td><td>1.0</td><td>2021-02-10T03:30:01Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187933943">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187933945">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jessie</td><td>1.0</td><td>2021-02-10T03:30:01Z</td><td>armel</td>
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
    <td>jessie</td><td>1.0</td><td>2021-02-10T03:30:01Z</td><td>armhf</td>
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
    <td>jessie</td><td>1.0</td><td>2021-02-10T03:30:01Z</td><td>i386</td>
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
    <td>jessie</td><td>1.0</td><td>2021-02-10T03:30:01Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187933971">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187933972">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.1</td><td>2021-02-10T03:30:02Z</td><td>amd64</td>
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
    <td>ascii</td><td>2.1</td><td>2021-02-10T03:30:02Z</td><td>arm64</td>
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
    <td>ascii</td><td>2.1</td><td>2021-02-10T03:30:02Z</td><td>armel</td>
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
    <td>ascii</td><td>2.1</td><td>2021-02-10T03:30:02Z</td><td>armhf</td>
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
    <td>ascii</td><td>2.1</td><td>2021-02-10T03:30:02Z</td><td>i386</td>
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
    <td>ascii</td><td>2.1</td><td>2021-02-10T03:30:02Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:30:03Z</td><td>amd64</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:30:03Z</td><td>arm64</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:30:03Z</td><td>armel</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:30:03Z</td><td>armhf</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:30:03Z</td><td>i386</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:30:03Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T20:35:45Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486266">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486272">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T20:35:45Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486276">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486283">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T20:35:45Z</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486286">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486304">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T20:35:45Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486308">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486313">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T20:35:45Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486323">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486329">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T20:35:45Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486333">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486339">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T20:37:49Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T20:37:49Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T20:37:49Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T20:37:49Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T20:37:49Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T20:37:49Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294563">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000206">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000210">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>amd64</td>
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
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>arm64</td>
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
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>armel</td>
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
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>armhf</td>
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
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000213">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187523523">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>i386</td>
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
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000224">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000226">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>m32</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000236">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000239">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000244">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000247">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000252">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000255">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000260">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000263">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>or1k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000266">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000271">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000275">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000278">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ascii-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ascii-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000280">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000282">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000291">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000294">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ascii</td><td>2.0.0</td><td>2021-02-10T03:27:13Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000297">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188000304">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742593">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742597">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>amd64</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>arm64</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>armel</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>armhf</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742602">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742610">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>i386</td>
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
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742618">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742622">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>m32</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742632">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742640">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742647">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742649">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742657">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742662">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742668">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742677">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>or1k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742687">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742697">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742703">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742709">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-beowulf-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-beowulf-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742713">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742720">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742729">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742737">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>beowulf</td><td>3.0</td><td>2021-02-10T03:27:13Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742740">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187742748">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486346">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486351">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486266">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486272">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486276">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486283">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486286">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486304">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486308">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486313">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486359">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486365">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486323">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486329">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486378">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486385">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>m32</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486391">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486394">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486402">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486410">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486421">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486428">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486437">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486447">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>or1k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486452">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486459">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486467">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486474">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486333">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486339">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486485">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486493">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486504">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486514">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>chimaera</td><td>4.0.0</td><td>2021-02-10T19:08:45Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486522">
            <img alt="errored" src="build-errored.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188486530">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294586">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294590">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>armel</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armel">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-armel.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armel-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-armel-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294600">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294610">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294621">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294631">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>m32</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294636">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294647">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294657">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294660">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294668">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294701">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294710">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294714">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>or1k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294720">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294727">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294734">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294740">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294563">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/devuan-ceres-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/devuan-ceres-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294753">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294758">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294770">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294776">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>ceres</td><td>1.0.0</td><td>2021-02-10T03:27:11Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294782">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188294811">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>alpha</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997400">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997402">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997404">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997405">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997406">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997408">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>armel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997410">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997417">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997419">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997422">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997424">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997425">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997426">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997429">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997431">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997432">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>m32</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997433">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997438">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>m68k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997439">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997440">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>mips</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997443">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997444">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>mipsel</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997445">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997449">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>or1k</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997453">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997455">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997456">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997457">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997458">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997461">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>s390</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997463">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997470">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997473">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997477">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>experimental</td><td>1.0.0</td><td>2021-02-10T03:27:05Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997479">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187997481">
            <img alt="failed" src="build-failed.svg" />
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
    <td>jessie</td><td></td><td>2021-02-10T16:27:00Z</td><td>armhf</td>
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
    <td>stretch</td><td></td><td>2021-02-10T16:44:18Z</td><td>armhf</td>
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
    <td>buster</td><td></td><td>2021-02-10T17:05:49Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-buster-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/raspbian-buster-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-buster-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/raspbian-buster-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bullseye</td><td></td><td>2021-02-10T17:34:01Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187935882">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/raspbian-bullseye-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/raspbian-bullseye-armhf-minbase.svg" />
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
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187739159">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187739169">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>2005-05-13T12:42:31Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187739174">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187739177">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>warty</td><td>4.10</td><td>2005-05-13T12:42:31Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187739180">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187739183">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187998998">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187998999">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187999001">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187999003">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187999007">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187999008">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187999015">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187999016">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hoary</td><td>5.04</td><td>2005-05-13T12:42:40Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187999019">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187999020">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993011">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993019">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993030">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993033">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>powerpc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993044">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993047">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993049">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993051">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993054">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993057">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>breezy</td><td>5.10</td><td>2005-10-13T19:34:42Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993060">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187993063">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932195">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932201">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932207">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932213">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-dapper-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-dapper-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932219">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932226">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932231">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>dapper</td><td>6.06</td><td>2006-05-31T18:59:06Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932235">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187932239">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705404">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705414">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705423">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705443">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705450">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705465">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-edgy-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-edgy-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705497">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>edgy</td><td>6.10</td><td>2006-10-25T17:07:09Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705502">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188705514">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722629">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722640">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722645">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>i386</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-i386">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-i386.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722664">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722679">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722686">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>powerpc</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-feisty-powerpc">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-feisty-powerpc.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722696">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>feisty</td><td>7.04</td><td>2007-04-17T18:14:27Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722702">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188722707">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>amd64</td>
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
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188085507">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188085513">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>i386</td>
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
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188085522">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188085525">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>lpia</td>
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
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>powerpc</td>
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
    <td>gutsy</td><td>7.10</td><td>2007-10-18T11:27:55Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188085546">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188085554">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>amd64</td>
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
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188711548">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188711555">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>i386</td>
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
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188711572">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188711583">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>lpia</td>
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
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>powerpc</td>
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
    <td>hardy</td><td>8.04</td><td>2008-09-20T01:13:43Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188711607">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188711611">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>amd64</td>
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
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188001661">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188001670">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>i386</td>
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
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188001685">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188001692">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>lpia</td>
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
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>powerpc</td>
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
    <td>intrepid</td><td>8.10</td><td>2008-11-19T21:01:09Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188001709">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188001712">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>amd64</td>
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
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>armel</td>
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
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>hppa</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187996113">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187996114">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>i386</td>
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
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187996120">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187996122">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>lpia</td>
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
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>powerpc</td>
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
    <td>jaunty</td><td>9.04</td><td>2009-04-22T21:35:16Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187996134">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187996138">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>amd64</td>
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
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>armel</td>
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
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>i386</td>
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
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187735947">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187735951">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>lpia</td>
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
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>powerpc</td>
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
    <td>karmic</td><td>9.10</td><td>2009-10-28T14:23:09Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187735972">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187735975">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>amd64</td>
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
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>armel</td>
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
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>i386</td>
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
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>ia64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188082257">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188082259">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>powerpc</td>
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
    <td>lucid</td><td>10.04</td><td>2010-04-29T17:24:55Z</td><td>sparc</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188082271">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188082280">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>maverick</td><td>10.10</td><td>2010-10-10T10:18:49Z</td><td>amd64</td>
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
    <td>maverick</td><td>10.10</td><td>2010-10-10T10:18:49Z</td><td>armel</td>
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
    <td>maverick</td><td>10.10</td><td>2010-10-10T10:18:49Z</td><td>i386</td>
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
    <td>maverick</td><td>10.10</td><td>2010-10-10T10:18:49Z</td><td>powerpc</td>
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
    <td>natty</td><td>11.04</td><td>2011-04-26T12:16:58Z</td><td>amd64</td>
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
    <td>natty</td><td>11.04</td><td>2011-04-26T12:16:58Z</td><td>armel</td>
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
    <td>natty</td><td>11.04</td><td>2011-04-26T12:16:58Z</td><td>i386</td>
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
    <td>natty</td><td>11.04</td><td>2011-04-26T12:16:58Z</td><td>powerpc</td>
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
    <td>oneiric</td><td>11.10</td><td>2011-10-12T18:19:26Z</td><td>amd64</td>
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
    <td>oneiric</td><td>11.10</td><td>2011-10-12T18:19:26Z</td><td>armel</td>
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
    <td>oneiric</td><td>11.10</td><td>2011-10-12T18:19:26Z</td><td>i386</td>
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
    <td>oneiric</td><td>11.10</td><td>2011-10-12T18:19:26Z</td><td>powerpc</td>
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
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>amd64</td>
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
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>armel</td>
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
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>armhf</td>
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
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>i386</td>
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
    <td>precise</td><td>12.04</td><td>2012-04-25T22:49:23Z</td><td>powerpc</td>
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
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>amd64</td>
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
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>arm64</td>
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
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>armhf</td>
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
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>i386</td>
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
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>powerpc</td>
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
    <td>trusty</td><td>14.04</td><td>2014-05-08T14:19:09Z</td><td>ppc64el</td>
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
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>amd64</td>
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
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>arm64</td>
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
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>armhf</td>
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
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>i386</td>
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
    <td>saucy</td><td>13.10</td><td>2014-05-08T14:19:34Z</td><td>powerpc</td>
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
    <td>raring</td><td>13.04</td><td>2014-05-08T14:19:55Z</td><td>amd64</td>
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
    <td>raring</td><td>13.04</td><td>2014-05-08T14:19:55Z</td><td>armhf</td>
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
    <td>raring</td><td>13.04</td><td>2014-05-08T14:19:55Z</td><td>i386</td>
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
    <td>raring</td><td>13.04</td><td>2014-05-08T14:19:55Z</td><td>powerpc</td>
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
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>amd64</td>
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
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>armel</td>
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
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>armhf</td>
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
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>i386</td>
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
    <td>quantal</td><td>12.10</td><td>2014-05-08T14:20:04Z</td><td>powerpc</td>
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
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>amd64</td>
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
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>arm64</td>
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
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>armhf</td>
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
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>i386</td>
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
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>powerpc</td>
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
    <td>utopic</td><td>14.10</td><td>2014-12-03T02:10:50Z</td><td>ppc64el</td>
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
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>amd64</td>
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
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>arm64</td>
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
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>armhf</td>
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
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>i386</td>
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
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>powerpc</td>
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
    <td>vivid</td><td>15.04</td><td>2015-04-24T18:46:12Z</td><td>ppc64el</td>
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
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>amd64</td>
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
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>arm64</td>
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
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>armhf</td>
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
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>i386</td>
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
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>powerpc</td>
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
    <td>wily</td><td>15.10</td><td>2015-10-22T12:47:57Z</td><td>ppc64el</td>
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
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>amd64</td>
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
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>arm64</td>
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
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>armhf</td>
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
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>i386</td>
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
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>powerpc</td>
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
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-xenial-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-xenial-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>xenial</td><td>16.04</td><td>2016-04-21T23:23:46Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187745849">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187745855">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>amd64</td>
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
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>arm64</td>
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
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>armhf</td>
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
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>i386</td>
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
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>powerpc</td>
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
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>yakkety</td><td>16.10</td><td>2016-10-13T13:26:23Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-yakkety-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-yakkety-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>amd64</td>
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
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>arm64</td>
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
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>armhf</td>
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
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>i386</td>
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
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>zesty</td><td>17.04</td><td>2017-04-13T13:44:04Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-zesty-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-zesty-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>amd64</td>
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
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188490884">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188490891">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188490895">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188490898">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>i386</td>
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
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188490920">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188490926">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>artful</td><td>17.10</td><td>2017-10-19T12:55:45Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188490931">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188490936">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>amd64</td>
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
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>arm64</td>
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
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>armhf</td>
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
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>i386</td>
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
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>bionic</td><td>18.04</td><td>2018-04-26T23:37:48Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-bionic-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-bionic-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>amd64</td>
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
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>arm64</td>
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
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188325964">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188325969">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>i386</td>
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
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>cosmic</td><td>18.10</td><td>2018-10-18T15:17:19Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-cosmic-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-cosmic-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>amd64</td>
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
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>arm64</td>
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
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>armhf</td>
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
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>i386</td>
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
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>disco</td><td>19.04</td><td>2019-04-18T09:40:05Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-disco-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-disco-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>amd64</td>
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
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>arm64</td>
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
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>armhf</td>
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
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>i386</td>
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
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>eoan</td><td>19.10</td><td>2019-10-30T17:43:41Z</td><td>s390x</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-s390x">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-s390x.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-eoan-s390x-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-eoan-s390x-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187724294">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-focal-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-focal-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>riscv64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187724349">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187724356">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>focal</td><td>20.04</td><td>2020-04-23T17:33:17Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187724368">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/187724376">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>amd64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-amd64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-amd64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-amd64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-amd64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>arm64</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-arm64">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-arm64.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-arm64-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-arm64-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>armhf</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-armhf">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-armhf.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-armhf-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-armhf-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188302787">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-i386-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-i386-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>ppc64el</td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-ppc64el">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-ppc64el.svg" />
        </a>
    </td>
    <td>
        <a href="https://hub.docker.com/r/lpenz/ubuntu-groovy-ppc64el-minbase">
            <img alt="passed" src="https://img.shields.io/docker/pulls/lpenz/ubuntu-groovy-ppc64el-minbase.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>riscv64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188302809">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188302813">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>groovy</td><td>20.10</td><td>2020-10-22T14:32:40Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188302818">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/188302823">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T22:04:19Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423228">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423231">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T22:04:19Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423235">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423241">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T22:04:19Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423254">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423259">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T22:04:19Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423264">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423272">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T22:04:19Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423276">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423281">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T22:04:19Z</td><td>riscv64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423286">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423292">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T22:04:19Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423299">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423305">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T21:34:13Z</td><td>amd64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423228">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423231">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T21:34:13Z</td><td>arm64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423235">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423241">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T21:34:13Z</td><td>armhf</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423254">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423259">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T21:34:13Z</td><td>i386</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423264">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423272">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T21:34:13Z</td><td>ppc64el</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423276">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423281">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T21:34:13Z</td><td>riscv64</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423286">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423292">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>
<tr>
    <td>hirsute</td><td>21.04</td><td>2021-02-10T21:34:13Z</td><td>s390x</td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423299">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
    <td>
        <a href="https://travis-ci.com/lpenz/docker-debian-releases/builds/192423305">
            <img alt="failed" src="build-failed.svg" />
        </a>
    </td>
</tr>

</tbody>
</table>
