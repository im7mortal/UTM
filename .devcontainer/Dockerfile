FROM golang:1.24-bookworm

# Match a typical host Linux UID/GID (override in devcontainer.json build args if needed).
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    git \
    sudo \
  && groupadd --gid "${USER_GID}" "${USERNAME}" \
  && useradd -s /bin/bash -m -u "${USER_UID}" -g "${USER_GID}" "${USERNAME}" \
  && mkdir -p /go/pkg /home/${USERNAME}/.cache/go-build \
  && chown -R "${USER_UID}:${USER_GID}" /go /home/${USERNAME} \
  && echo "${USERNAME} ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/${USERNAME} \
  && chmod 440 /etc/sudoers.d/${USERNAME} \
  && rm -rf /var/lib/apt/lists/*

ENV GOPATH=/go
ENV PATH=${GOPATH}/bin:/usr/local/go/bin:${PATH}

USER ${USERNAME}


