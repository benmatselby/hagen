# Hagen

![Go](https://github.com/benmatselby/hagen/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/benmatselby/hagen?style=flat-square)](https://goreportcard.com/report/github.com/benmatselby/hagen)

![Tom Hagen](/img/hagen.jpg)

CLI application for getting information out of [GitHub](https://github.com), mainly for running a sprint and generating a burn down and creating changelogs.

```text
CLI application for retrieving data from GitHub

Usage:
  hagen [command]

Available Commands:
  burndown    Provide a burndown table for a project
  help        Help about any command
  issues      List issues given the search criteria. Default query is to list issues where the author is ${GITHUB_OWNER}
  projects    List the projects. Default query is to list project for ${GITHUB_OWNER}
  repos       List the repositories based on a query. Default query is to list repos by ${GITHUB_OWNER}

Flags:
      --config string   config file (default is $HOME/.benmatselby/hagen.yaml)
  -h, --help            help for hagen

Use "hagen [command] --help" for more information about a command.
```

## Requirements

- [Go version 1.13+](https://golang.org/dl/)
- You will need require a [PAT from GitHub](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line).

## Configuration

There are two aspects to configuring `hagen`.

### Environment variables

In order to connect to GitHub `hagen` requires three environment variables.

```bash
export GITHUB_OWNER=""  # This is likely to be your personal GitHub username
export GITHUB_ORG=""    # This is if you are linked to a GitHub org
export GITHUB_TOKEN=""  # Your PAT
```

The token can be generated by reading this [article](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line). Typically the `GITHUB_OWNER` will be your username.

### Configuration file

`hagen` tries to save on typing, so you can define some queries in a configuration file.

```yml
templates:
  sprint:
    query: "state:open project:acme/14"
    count: 200
  open-bugs:
    query: "repo:benmatselby/donny state:open"
    count: 50
```

To understand how to write the `query` section, follow [this article](https://help.github.com/en/articles/searching-issues-and-pull-requests).

## Installation via Docker

Other than requiring [docker](http://docker.com) to be installed, there are no other requirements to run the application this way. This is the preferred method of running the `hagen`. The image is [here](https://hub.docker.com/r/benmatselby/hagen/).

```shell
$ docker run \
  --rm \
  -t \
  -eGITHUB_ORG \
  -eGITHUB_OWNER \
  -eGITHUB_TOKEN \
  -v "${HOME}/.benmatselby":/root/.benmatselby \
  benmatselby/hagen:latest "$@"
```

The `latest` tag mentioned above can be changed to a released version. For all releases, see [here](https://hub.docker.com/repository/docker/benmatselby/hagen/tags). An example would then be:

```shell
benmatselby/hagen:version-2.0.0
```

This would use the `verson-2.2.0` release in the docker command.

## Installation via Git

```bash
git clone git@github.com:benmatselby/hagen.git
cd hagen
make all
./hagen
```

You can also install into your `$GOPATH/bin` by `go install`
