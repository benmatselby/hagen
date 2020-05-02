# Changelog

## next

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
