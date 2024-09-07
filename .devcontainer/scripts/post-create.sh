#!/bin/sh -e

# Setup zsh plugins
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

# switch to zsh
#chsh -s $(which zsh)
sudo chsh -s $(which zsh)

# Setup plandex
#curl -sL https://plandex.ai/install.sh | bash

echo "done 4" \
    && apt-get update \
    && apt-get upgrade -y
