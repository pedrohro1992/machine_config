#!/bin/bash

#List all tmux sessions open a make a fuzzy finder on the list

TMUX_SESSION=$(tmuxifier list-sessions | fzf --reverse)

#Uses tmuxifier to load that session

tmuxifier edit-session $TMUX_SESSION