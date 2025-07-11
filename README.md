# Docker Go SDK Examples

This repository contains examples of using the [Docker Go SDK](https://github.com/docker/go-sdk).

## Running the examples

```shell
go test -v -timeout 30s -count=1 ./...
```

## Container Runtime support

These examples are tested on GithHub runners with the following container runtimes:

| Container Runtime    | Worker | Status |
| -------------------- | ------ | ------ |
| Docker               | Ubuntu | ✅ |
| Docker (Rootless)    | Ubuntu | ✅ |
| Docker (Containerd)  | Ubuntu | ✅ |
| Podman               | Ubuntu | ✅ |
| Colima               | macOS  | ❌ (pending) |
| Rancher Desktop      | macOS  | ❌ (pending) |
| Orbstack             | macOS  | ❌ (pending) |

The GitHub runners are running on Ubuntu, using the `ubuntu-latest` runner.

For further details, see the [GitHub Actions workflow](.github/workflows/ci-test-go.yml).

> [!NOTE]
> The container runtimes using macOS runners are not supported yet because of the lack of support for the nested virtualization.
>
> For further details, see this [GitHub issue](https://github.com/docker/setup-docker-action/pull/53#issuecomment-1923467713).
