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

