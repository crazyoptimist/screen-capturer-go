#!/bin/sh
# Abort on error
set -e

EXE="/usr/local/bin/screen-server-darwin"
REG="/Library/LaunchAgents/net.crazyoptimist.screen-server-darwin.plist"

printf 'Have you replaced the vhost name with yours in the service file (y/n)? '
read answer

if [ "$answer" != "${answer#[Yy]}" ] ;then 
  sudo chmod +x ./screen-server-darwin
  sudo cp ./screen-server-darwin $EXE
  sudo cp ./net.crazyoptimist.screen-server-darwin.plist $REG

  sudo launchctl load $REG
  echo "Successfully installed the application."
else
  echo "Exiting..."
  exit
fi
