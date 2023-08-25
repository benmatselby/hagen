# Changelog

## next

## 3.5.0

- Bump Go version to 1.21
- Allow the caller of `make docker-build` to specify the Docker platform.

## 3.4.0

- Push Docker images to GitHub Package Registry.

## 3.3.0

- Bumped docker image to Go 1.17 runtime.

## 3.2.0

- Bumped docker image to Go 1.16 runtime.
- Bump the build environment to test on 1.16.

## 3.1.0

- Addition of a `--labels` option for the issues command. [#32](https://github.com/benmatselby/hagen/pull/32)
- Do not ask "Press enter for more results" if there are no more results to show.
- Bump dependencies

## 3.0.0

- Bumped docker image to Go 1.14 runtime.
- Bump the build environment to test on 1.12, 1.13, and 1.14.
- Restructure the commands to limit the typing.
  - Rationale is to make this like [walter](http://github.com/benmatselby/walter), and [lionel](http://github.com/benmatselby/lionel)
- Provide a `projects` command to list projects.
- Provide a `burndown` command.
- Provide a "--recursive" flag for `repo ls` so it will get all results for you. If not, it will ask you to "page" with the enter key.

## 2.0.0

- Standardise on command naming.
- Dockerised version configured and deployed.

## 1.1.0

- New command to list repos based on a query `repo ls`.

## 1.0.0

- First versioned release.
