#!/bin/sh

EXE="/usr/local/bin/screen-server-darwin"
REG="/Library/LaunchAgents/net.crazyoptimist.screen-server-darwin.plist"

sudo launchctl unload -F $REG

sudo killall screen-server-darwin

sudo rm -f $EXE $REG

echo "Uninstalled the application successfully."
