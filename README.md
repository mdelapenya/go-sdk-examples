# Docker Go SDK Examples

This repository contains examples of using the [Docker Go SDK](https://github.com/docker/go-sdk).

## Running the examples

```shell
go test -v -timeout 30s -count=1 ./...
```

## Container Runtime support

These examples are tested on GithHub runners with the following container runtimes:

- Docker
- Podman

The GitHub runners are running on Ubuntu 22.04.

For further details, see the [GitHub Actions workflow](.github/workflows/ci-test-go.yml).
