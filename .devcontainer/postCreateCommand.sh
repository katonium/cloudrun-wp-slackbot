#!/bin/bash
# 
# This script installs the necessary dependencies for the project
# This script is executed after the DevContainer is created

# Install mise
# https://github.com/jdx/mise
curl https://mise.run | sh

# Install Claude Code
# https://docs.anthropic.com/en/docs/claude-code/setup
npm install -g @anthropic-ai/claude-code

# Install CloudNative Buildpacks
# https://buildpacks.io/docs/for-platform-operators/how-to/integrate-ci/pack/
go get -u github.com/buildpacks/pack    

cd /home/codespace/ && \
    curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-linux-x86_64.tar.gz && \
    tar -xf google-cloud-cli-linux-x86_64.tar.gz  
# Run /home/codespace/google-cloud-sdk/install.sh to complete the installation
