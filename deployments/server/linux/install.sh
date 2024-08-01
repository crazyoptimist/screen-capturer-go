#!/bin/sh
# Abort on error
set -e

printf 'Have you replaced the vhost name with yours in the service file (y/n)? '
read answer

if [ "$answer" != "${answer#[Yy]}" ] ;then 
  sudo chmod +x ./screen-server-linux
  sudo cp ./screen-server-linux /usr/local/bin/
  sudo cp ./screen-capturer.service /etc/systemd/system/

  echo "Successfully installed the application."
else
  echo "Exiting..."
  exit
fi
