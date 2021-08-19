[![CI](https://github.com/lpenz/docker-debian-releases/actions/workflows/ci.yml/badge.svg)](https://github.com/lpenz/docker-debian-releases/actions/workflows/ci.yml)


# docker-debian-releases

This repository creates docker images of Debian-based system using
debootstrap, for various architectures, and uploads them
to [docker hub](https://hub.docker.com/r/lpenz/) using github actions.

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
- [branch-info](branch-info): python script that translates the name
  of the branch being pushed to all the information about the image to
  be built and deployed.
- [dockerhub-set-descriptions](go/cmd/dockerhub-set-descriptions/main.go):
  updates the short and long description of the image in docker hub
  after a new version is deployed.


The [status page](http://www.lpenz.org/docker-debian-releases/) is
also created by a github action, from information obtained from
scrapping all available mirrors. The following scripts are in
charge of this process:
- [apt-mirror-info](go/cmd/apt-mirror-info/main.go): scraps Debian
  and Ubuntu repositories and outputs a json with information about
  all releases it can find.
- [json-tmpl-render](go/cmd/json-tmpl-render/main.go): renders a
  template file with information from a json file.
- [README.md.tmpl](README.md.tmpl): template for README.md that uses
  the information obtained by apt-mirror-info and travis-branch-jobs
  to create a table of images and status' with links to jobs.
- [SConstruct](SConstruct): scons script that builds the go sources
  and README.md.
