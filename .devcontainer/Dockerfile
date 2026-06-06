FROM ghcr.io/lugolven/bazelisk-ubuntu-container:1.29.0-ubuntu25.10

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    gcc \
    git \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

ARG GO_VERSION=1.24.4
RUN ARCH=$(dpkg --print-architecture) && \
    curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz | tar -C /usr/local -xz
ENV PATH="/usr/local/go/bin:${PATH}"

RUN useradd -m -s /bin/bash devcontainer
USER devcontainer
