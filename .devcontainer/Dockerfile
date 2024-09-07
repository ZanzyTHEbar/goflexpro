FROM mcr.microsoft.com/devcontainers/go:latest

ARG DEBIAN_FRONTEND=noninteractive

RUN echo "Updating Distro" \
    && apt-get update \
    && apt-get upgrade -y

RUN DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y build-essential make \
    ca-certificates \
    git \
    libssl-dev \
    zlib1g-dev \
    libbz2-dev \
    libreadline-dev \
    libsqlite3-dev \
    wget \
    curl \
    llvm \
    libncurses5-dev \
    xz-utils \
    tk-dev \
    libxml2-dev \
    libxmlsec1-dev \
    libffi-dev \
    liblzma-dev \
    autojump \
    zsh \
    fonts-powerline

# Setup oh-my-zsh
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.5/zsh-in-docker.sh)" -- \
    -p git -p autojump -p zsh-autosuggestions -p zsh-syntax-highlighting

# Install Go tools





