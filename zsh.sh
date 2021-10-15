#!/bin/bash
#INSTALL ZSH
sudo apt install zsh

#CONFIGURE Oh My Zsh

curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh; zsh


#Install Plugins
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting 

git clone https://github.com/zsh-users/zsh-autosuggestions $ZSH_CUSTOM/plugins/zsh-autosuggestions

git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf && ~/.fzf/install

git clone https://github.com/asdf-vm/asdf.git ~/.asdf

#Copy ZSH FIle

wget https://raw.githubusercontent.com/pedrohro1992/machine_config/main/zshrc
mv zshrc ~/.zshrc 


#Stern ASDF Plugins

asdf plugin add  kubectl 
asdf plugin add terraform
asdf plugin add golang
asdf plugin add kops
asdf plugin add python
asdf plugin add stern 


#Install All the Latest Version


asdf install kubectl latest
asdf install terraform latest
asdf install golang latest
asdf install kops latest
asdf install python latest
asdf install stern latest


#Set The global off all plugins to latest

#asdf global kubectl latest
asdf global terraform latest
asdf global golang latest
asdf global kops latest
asdf global python latest
asdf global stern latest
