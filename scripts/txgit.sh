#!/bin/bash
export EDITOR=code
echo "Git repo to clone: "
read GIT_REPO
if [[ $GIT_REPO =~ .*gitlab.* ]]; then
  REPO_NAME=$(basename $GIT_REPO | awk '{ print substr( $0, 1, length($0)-4 ) }')
  git clone $GIT_REPO /home/pedro/Workspace/work/$REPO_NAME
  tmuxifier new-session $REPO_NAME
else 
  REPO_NAME=$(basename $GIT_REPO | awk '{ print substr( $0, 1, length($0)-4 ) }')
  git clone $GIT_REPO /home/pedro/Workspace/personal/$REPO_NAME
  tmuxifier new-session $REPO_NAME   
fi
#Set a variable to get the basename of the repo


echo "$REPO_NAME"