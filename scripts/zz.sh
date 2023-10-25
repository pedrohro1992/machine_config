#!/bin/bash

#DEFINE THE DIR WITCH TMUX WILL STAR THE SEASON
TMUX_DIR=$(zoxide query -l | fzf --reverse)

#START NEW TMUX SESSION ON DIR SELECTED VIA FUZZY FINDER
tmux new-session -s teste -c $TMUX_DIR