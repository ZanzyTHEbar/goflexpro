if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi

autoload -U colors && colors
PS1="%B%{$fg[red]%}[%{$fg[yellow]%}%n%{$fg[green]%}@%{$fg[blue]%}%M %{$fg[magenta]%}%~%{$fg[red]%}]%{$reset_color%}$%b "

#autoload -U compinit && compinit
#eval "$(register-python-argcomplete pipx)"

export ZSH="/root/.oh-my-zsh"

plugins=(git zsh-autosuggestions zsh-syntax-highlighting autojump poetry)

HISTSIZE=10000
SAVEHIST=10000
setopt appendhistory

# Basic auto/tab complete:
autoload -U compinit
zstyle ':completion:*' menu select
zmodload zsh/complist
compinit
_comp_options+=(globdots)               # Include hidden files.

# Custom ZSH Binds
bindkey '^ ' autosuggest-accept

# Load aliases and shortcuts if existent.
[ -f "$HOME/zsh/aliasrc" ] && source "$HOME/zsh/aliasrc"

# Load ; should be last.
source ~/.oh-my-zsh/custom/themes/powerlevel10k/powerlevel10k.zsh-theme
source $ZSH/oh-my-zsh.sh

# To customize prompt, run `p10k configure` or edit ~/zsh/.p10k.zsh
[[ ! -f ~/zsh/.p10k.zsh ]] || source ~/zsh/.p10k.zsh
# comment this out if you want to use the wizard
POWERLEVEL9K_DISABLE_CONFIGURATION_WIZARD=true

# Created by `pipx` on 2024-05-04 20:06:06
export PATH="$PATH:~/.local/bin"
export PATH="$PATH:/root/.local/bin"
export LANG='en_US.UTF-8'
export LANGUAGE='en_US:en'
export LC_ALL='en_US.UTF-8'
export TERM=xterm

POWERLEVEL9K_SHORTEN_STRATEGY="truncate_to_last"
POWERLEVEL9K_LEFT_PROMPT_ELEMENTS=(user dir vcs status)
POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=()
POWERLEVEL9K_STATUS_OK=false
POWERLEVEL9K_STATUS_CROSS=true

DISABLE_AUTO_UPDATE=true
DISABLE_UPDATE_PROMPT=true