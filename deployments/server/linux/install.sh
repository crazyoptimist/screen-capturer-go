#!/bin/sh
# Abort on error
set -e

EXE="/usr/local/bin/screen-server-linux"
REG="${HOME}/.config/systemd/user"

printf 'Have you replaced the vhost name with yours in the service file (y/n)? '
read answer

if [ "$answer" != "${answer#[Yy]}" ] ;then 
  sudo chmod +x ./screen-server-linux
  sudo cp ./screen-server-linux $EXE
  mkdir -p $REG
  cp ./screen-capturer.service "$REG/"
  systemctl --user enable screen-capturer.service
  systemctl --user start screen-capturer.service

  echo "Successfully installed the application."
else
  echo "Exiting..."
  exit
fi
