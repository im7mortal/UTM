#!/usr/bin/env bash
set -euo pipefail

if [ -f "go.mod" ]; then
  go mod tidy
fi

# Verify the container toolchain and project health immediately after creation.
go version

