# GitHub Copilot Instructions for hagen

## Project Overview

- **hagen** is a CLI tool for retrieving and displaying GitHub issues, pull requests, and repositories.
- Written in Go, using Cobra for CLI, Google go-github for API, and tablewriter for table output.

## Coding Guidelines

- Follow idiomatic Go practices and use `gofmt` for formatting.
- Use the strategy pattern for output formatting (see `cmd/list_issues.go`).
- Write tests for all new features and bug fixes (see `cmd/list_issues_test.go`).
- Use dependency injection for clients (see how `hagen.Provider` is used).
- Prefer clear, user-friendly CLI flags and help text.
- Use semantic commit messages.
- Ensure all code passes linting (`make lint`).

## Features & Conventions

- Issue listing supports both default and table output (`--display` flag).
- Table output includes status, and human-readable dates (`--human-dates`).
- All user-facing strings should be clear and concise.
- Keep the codebase well-documented and update the `CHANGELOG.md` for every release.

## Testing

- Use `make test-cov` to run all tests.
- Add tests for new strategies, flags, and output formats.
- Use table-driven tests where appropriate.
- Where applicable use an io.Writer that captures the output and then compare it against expected output. (see `cmd/list_issues_test.go` for examples).

## Dependencies

- Use `go get` to add dependencies.
- Keep dependencies up to date and tidy with `go mod tidy`.

## Pull Requests

- Ensure all tests pass before submitting a PR.
- Update documentation and changelog as needed.
- Keep PRs focused and well-described.

## Misc

- Use the `main.go` entrypoint for CLI execution.
- Dockerfile and Makefile are provided for builds and CI.

## Changelog

- When asked to create a release, use the last git tag and summarise the commits since.
- Increment the version in the `CHANGELOG.md` file and define a new heading.
- Follow the format in `CHANGELOG.md` for consistency, for example:
  - `## next` for the next version.
  - `## x.y.z` for previous versions, with bullet points for changes.
  - Do not use the date in the changelog, just the version number.
