
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
source "/home/$USER/Workspace/Misc/machine_config/kube-ps1/kube-ps1.sh"
[ -f ~/Workspace/Misc/machine_config/kubectl_alias/.kubectl_aliases  ] && source ~/Workspace/Misc/machine_config/kubectl_alias/.kubectl_aliases 

#PS1 CONFIGURATIONS
PROMPT=$PROMPT'$(kube_ps1)'

#TURNOFF KUBE_PS1 BY DEFAULT
kubeoff


#KUBERNETS CONTEXT ALIAS
alias kdev='export KUBECONFIG=~/.kube/develop.yaml'
alias kqa='export KUBECONFIG=~/.kube/qa.yaml'
alias khml='export KUBECONFIG=~/.kube/homolog.yaml'
alias khml2='export KUBECONFIG=~/.kube/homolog-02.yaml'
alias khml3='export KUBECONFIG=~/.kube/homolog-03.yaml'
alias kpixdr='export KUBECONFIG=~/.kube/pix-genial-dr.yaml'
alias kpreprd='export KUBECONFIG=~/.kube/preprod.yaml'
alias kparceiros='export KUBECONFIG=~/.kube/parceiros.yaml'
alias kprod="export KUBECONFIG=/home/pedro/.kube/production.yaml"
alias kpreprd2="export KUBECONFIG=/home/pedro/.kube/preprod-02.yaml"
alias kpixprd="export KUBECONFIG=/home/pedro/.kube/pix-genial-production.yaml"
alias kpixhml="export KUBECONFIG=/home/pedro/.kube/pix-genial-homolog.yaml"
alias tinit='terraform init'
alias tplan='terraform plan'
alias tapply='terraform apply'
alias tdev='terraform workspace select develop'
alias thml='terraform workspace select homolog'
alias tqa='terraform workspace select qa'
alias tprd='terraform workspace select production'
alias tpreprd='terraform workspace select preprod'
alias trplandebug='terragrunt plan --terragrunt-log-level debug --terragrunt-debug'
alias trapplydebug='terragrunt apply --terragrunt-log-level debug --terragrunt-debug'

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh


export GOPATH=/home/pedro/Workspace/go
