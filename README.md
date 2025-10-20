# Docker Go SDK Examples

This repository contains examples of using the [Docker Go SDK](https://github.com/docker/go-sdk).

## Running the examples

```shell
go test -v -timeout 30s -count=1 ./...
```

## Container Runtime support

These examples are tested on GitHub runners with the following container runtimes:

| Container Runtime    | Worker | Status |
| -------------------- | ------ | ------ |
| Docker               | Ubuntu | ✅ |
| Docker (Rootless)    | Ubuntu | ✅ |
| Docker (Containerd)  | Ubuntu | ✅ |
| Docker (Lima)        | macOS Intel (15) | ✅ |
| Podman               | Ubuntu | ✅ |
| Podman               | macOS | ✅ |
| Colima (Docker)      | macOS Intel (15) | ✅ |
| Colima (Containerd)  | macOS Intel (15) | ✅ |
| Rancher Desktop      | macOS  | ❌ (pending) |
| Orbstack             | macOS  | ❌ (not supported - nested virtualization required) |

The GitHub runners are configured as follows:
- **Ubuntu runners**: Using `ubuntu-latest` with native Docker support
- **macOS Intel runners**: Using `macos-15-intel` with Docker via Lima VM (official `docker/setup-docker-action`)

For further details, see the [GitHub Actions workflow](.github/workflows/ci-test-go.yml).

> [!NOTE]
> macOS ARM (Apple Silicon) runners are not currently used due to lack of nested virtualization support on M1/M2 chips.
> Intel-based macOS runners (`macos-15-intel`) will be supported by GitHub until Fall 2027.
>
> For further details, see this [GitHub issue](https://github.com/docker/setup-docker-action/pull/53#issuecomment-1923467713).
