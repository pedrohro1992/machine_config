#!/bin/bash
#INSTALL ZSH
sudo apt install zsh

#CONFIGURE Oh My Zsh

curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh; zsh

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

#Copy ZSH Source Configurations


cp zshrc ~/.zshrc 

#Copy history files
cp zsh_history ~/.zsh_history 


