FROM ubuntu:22.04

ARG GITHUB_RUNNER_VERSION="2.300.0"

ENV GITHUB_OWNER "myorganization"
ENV RUNNER_WORKDIR "_work"
ENV TZ="Europe/London"

ARG DEBIAN_FRONTEND="noninteractive"


RUN apt-get update \
  && apt-get install -y \
  ca-certificates \
  curl \
  apt-transport-https \
  lsb-release \
  gnupg \
  && curl -sL https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor | tee /etc/apt/trusted.gpg.d/microsoft.gpg > /dev/null \
  && AZ_REPO=$(lsb_release -cs) \
  && echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ $AZ_REPO main" | tee /etc/apt/sources.list.d/azure-cli.list \
  && apt-get update \
  && apt-get install -y \
  azure-cli \
  iputils-ping \
  sudo \
  git \
  unzip \
  jq \
  gh

# Required by "hashicorp/setup-terraform"
RUN curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash - \
  && sudo apt-get install -y nodejs
RUN apt-get clean \
  && rm -rf /var/lib/apt/lists/*

RUN adduser --uid 1000 --gecos "GitHub Runner" --disabled-password github-runner && \
  echo 'github-runner ALL=(ALL) NOPASSWD:ALL' | sudo EDITOR='tee -a' visudo
USER 1000

WORKDIR /home/github-runner

# Install everything needed for the GitHub Action self-hosted-runner
RUN curl -Ls https://github.com/actions/runner/releases/download/v${GITHUB_RUNNER_VERSION}/actions-runner-linux-x64-${GITHUB_RUNNER_VERSION}.tar.gz | tar xz
RUN sudo ./bin/installdependencies.sh

COPY ./scripts/*.sh /home/github-runner/scripts/
RUN sudo chmod +x /home/github-runner/scripts/*.sh

ENV PATH="${PATH}:/home/github-runner/scripts"
