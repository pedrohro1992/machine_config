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
source "/home/$USER/Workspace/personal/machine_config/kube-ps1/kube-ps1.sh"
[ -f ~/Workspace/personal/machine_config/kubectl_alias/.kubectl_aliases  ] && source ~/Workspace/personal/machine_config/kubectl_alias/.kubectl_aliases 

#PS1 CONFIGURATIONS
PROMPT=$PROMPT'$(kube_ps1)'

#alias Kuberntes
alias kclear='kubectx -u'
alias kctx='kubectx'
alias kns='kubens'

#alias Terraform
alias tinit='terraform init'
alias tplan='terraform plan'
alias tapply='terraform apply'
alias tapplyforce='terraform apply --auto-approve'
alias tplantg='terraform plan --target='
alias tapplytg='terraform apply --target='

#Alias neo vim 
alias nvim='/usr/bin/nvim.appimage' 

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh


eval "$(starship init zsh)"
