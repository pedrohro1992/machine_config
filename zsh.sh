#!/bin/bash
#INSTALL ZSH
sudo apt install zsh

#CONFIGURE Oh My Zsh

curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh; zsh

sudo usermod --shell $(which zsh) $USER


#Install Plugins
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting 

git clone https://github.com/zsh-users/zsh-autosuggestions $ZSH_CUSTOM/plugins/zsh-autosuggestions

git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf && ~/.fzf/install

git clone https://github.com/asdf-vm/asdf.git ~/.asdf

#Copy ZSH FIle

wget https://raw.githubusercontent.com/pedrohro1992/machine_config/main/zshrc
mv .zshrc ~/.zshrc 


