
export ZSH="$HOME/.oh-my-zsh"


ZSH_THEME="robbyrussell"

plugins=(git
         dnf
         zsh-syntax-highlighting
         zsh-autosuggestions
         asdf
        )

#STARTUPS SCRIPTS
source $ZSH/oh-my-zsh.sh
source "/home/$USER/Workspace/machine_config/kube-ps1/kube-ps1.sh"
[ -f ~/Workspace/machine_config/kubectl_alias/.kubectl_aliases  ] && source ~/Workspace/machine_config/kubectl_alias/.kubectl_aliases 

#PS1 CONFIGURATIONS
PROMPT=$PROMPT'$(kube_ps1)'

#TURNOFF KUBE_PS1 BY DEFAULT
kubeoff


#KUBERNETS ALIAS
alias off='shutdown -h now'
alias forceoff='systemctl poweroff -i'
alias amatec='sudo /usr/bin/python3 /home/ifarted/Workspace/G213Colors/G213Colors.py -c FF8C00' 
[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

#GOPATH
export GOPATH=/home/ifarted/Workspace/golang/go

### MANAGED BY RANCHER DESKTOP START (DO NOT EDIT)
export PATH="/home/ifarted/.rd/bin:$PATH"
### MANAGED BY RANCHER DESKTOP END (DO NOT EDIT)
