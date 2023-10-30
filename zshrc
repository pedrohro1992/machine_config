export ZSH="$HOME/.oh-my-zsh"
export PATH="$HOME/.tmuxifier/bin:$PATH"
export EDITOR=code

ZSH_THEME="robbyrussell"

plugins=(git
         dnf
         zsh-syntax-highlighting
         zsh-autosuggestions
         asdf
        )

#STARTUPS SCRIPTS
source $ZSH/oh-my-zsh.sh
#source "/home/$USER/Workspace/personal/machine_config/kube-ps1/kube-ps1.sh"
[ -f ~/Workspace/personal/machine_config/kubectl_alias/.kubectl_aliases  ] && source ~/Workspace/personal/machine_config/kubectl_alias/.kubectl_aliases 

#PS1 CONFIGURATIONS
#PROMPT=$PROMPT'$(kube_ps1)'


#Misc Alias
alias goto='DIR=$(zoxide query -l | fzf --reverse --print0) ; cd $DIR'
alias fcode='/home/pedro_hro/Workspace/personal/machine_config/scripts/fcode.sh'

# tmuxifier alias 
alias txls='/home/pedro_hro/Workspace/personal/machine_config/scripts/txls.sh'
alias txedit='/home/pedro_hro/Workspace/personal/machine_config/scripts/txedit.sh'

#Terraform alias
alias t='terraform'
alias tinit='terraform init'
alias tplan='terraform plan'
alias tapply='terraform apply'

#Terragrunt alias
alias tr='terragrunt'
alias trplan='terragrunt plan'
alias trapply='terragrunt apply'

#Kubectx Alias
alias kunset='kubectx --unset'

#zoxide
alias cz='z'

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh

eval "$(starship init zsh)"

eval "$(zoxide init zsh)"

eval "$(tmuxifier init -)"


#export GOPATH=/home/pedro/workspace/go
