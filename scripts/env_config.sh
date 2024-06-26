#!/bin/bash
#INSTALL ZSH
sudo apt install zsh

#CONFIGURE Oh My Zsh

curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh
zsh

sudo usermod --shell $(which zsh) $USER

#Install AWS CLI

#TODO Add verification of the AWS CLI V1 installation

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

#Install Plugins
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

git clone https://github.com/zsh-users/zsh-autosuggestions $ZSH_CUSTOM/plugins/zsh-autosuggestions

git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf && ~/.fzf/install

git clone https://github.com/asdf-vm/asdf.git ~/.asdf

git clone https://github.com/jeffreytse/zsh-vi-mode $ZSH_CUSTOM/plugins/zsh-vi-mode

#Copy ZSH Source Configurations

cp dotfiles/zsh/zshrc ~/.zshrc


#Install Treesetier dependecies
sudo dnf install g++
