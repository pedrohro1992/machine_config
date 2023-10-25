#!/bin/bash

CODE_DIR=$(zoxide query -l | fzf --reverse)

code $CODE_DIR