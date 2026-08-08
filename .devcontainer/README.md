# Dev Container Notes

This project uses a custom Go dev container built from `.devcontainer/Dockerfile`.

## What it sets up

- Stable Go toolchain in an isolated container.
- Persistent caches for Go modules and build artifacts.
- Basic editor defaults for formatting and linting.
- Automatic `go mod tidy` and container-local Git setup on first container create.
- A Dockerfile-based setup that avoids the feature-image and UID-update temp-file failures seen in this environment.

## Rebuild tips

If dependencies or base image behavior changes, rebuild the container from your IDE.


