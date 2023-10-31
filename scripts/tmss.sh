#!/bin/bash

#Define the seasson to go 

SESSION=$(tmux list-sessions -F '#{session_name}' | fzf --reverse)

#Connet to session 

tmux attach-session -t $SESSION
