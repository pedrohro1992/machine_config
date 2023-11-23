#!/bin/bash

# Find the window IDs of Terminal and Firefox
terminal_window=$(xdotool search --onlyvisible --class WezTerm)
firefox_window=$(xdotool search --onlyvisible --class firefox)
# Minimize all windows except Terminal and Firefox
for window in $(xdotool search --onlyvisible ".*" | grep -vE "$terminal_window|$firefox_window")
do
    xdotool windowminimize "$window"
done

