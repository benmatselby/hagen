# Changelog

## Release-next

- Bump Go version to 1.26

## 3.14.0

- Bump golang base image to 1.25.0-alpine
- Provide a "Mean time to merge" display strategy for `issues` command.

## 3.13.0

- Add `templates` command to list all templates from the config file, with `-v` to show queries
- Bump golang base image to 1.24.4-alpine
- Bump github.com/olekukonko/tablewriter to 1.0.7

## 3.12.0

- Display the timezone in issues list

## 3.11.0

- Add Copilot instructions for the project
- Updated dependencies
- Provide more examples for the `issues list` command

## 3.10.0

- Add strategy pattern for displaying issues, with a new `--display` flag to choose between default and table output.
- Add table output strategy using `tablewriter`, with columns for Type (Issue/Pull Request), Repository, Number, Title, Labels, Status, Created At, and Closed At.
- Add `--human-dates` flag to display dates in a human-readable format (e.g., "Monday 02 January, 2006 at 15:04").
- Add Status and Closed At columns to the table output.
- Refactor and expand tests to cover all new display options and output formats.

## 3.9.0

- Bump Go version to 1.24

## 3.8.0

- Bump Go version to 1.22

## 3.7.0

- Add in the release version to the auto-generated binaries during a release.

## 3.6.0

- Addition of binaries for each release, documented on the release page.

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

## 2.0.0

- Standardise on command naming.
- Dockerised version configured and deployed.

## 1.1.0

- New command to list repos based on a query `repo ls`.

## 1.0.0

- First versioned release.
