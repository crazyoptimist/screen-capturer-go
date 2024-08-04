#!/bin/sh
# Abort on error
set -e

EXE="/usr/local/bin/screen-server-linux"

if [ ${XDG_SESSION_TYPE} = "x11" ] ;then 
  sudo chmod +x ./screen-server-linux
  sudo cp ./screen-server-linux $EXE

  echo "Successfully installed the application."
else
  echo "Your GUI is not Xorg, or active display is not found."
  exit
fi
