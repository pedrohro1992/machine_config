
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
source "/home/$USER/workspace/misc/machine_config/kube-ps1/kube-ps1.sh"
[ -f ~/workspace/misc/machine_config/kubectl_alias/.kubectl_aliases  ] && source ~/workspace/misc/machine_config/kubectl_alias/.kubectl_aliases 

#PS1 CONFIGURATIONS
PROMPT=$PROMPT'$(kube_ps1)'

#TURNOFF KUBE_PS1 BY DEFAULT
kubeoff



[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh


export GOPATH=/home/pedro/workspace/go
