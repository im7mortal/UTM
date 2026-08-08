#!/usr/bin/env bash
#
# Container-only git identity.
#
# Keep container git config isolated from the host so commits can be created
# without depending on host signing or identity configuration.

set -euo pipefail

GIT_NAME="${CONTAINER_GIT_NAME:-Devcontainer}"
GIT_EMAIL="${CONTAINER_GIT_EMAIL:-devcontainer@local}"

git config --global user.name "$GIT_NAME"
git config --global user.email "$GIT_EMAIL"
git config --global commit.gpgsign false
git config --global tag.gpgsign false
git config --global --add safe.directory "${PWD}"

echo "Container git identity configured (global, container-only):"
git config --global --show-origin user.name
git config --global --show-origin user.email
git config --global --show-origin commit.gpgsign

