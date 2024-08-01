#!/bin/sh

EXE="/usr/local/bin/screen-server-linux"
REG="${HOME}/.config/systemd/user"

systemctl --user stop screen-capturer.service
systemctl --user disable screen-capturer.service

sudo rm -f "$REG/screen-capturer.service" $EXE

echo "Uninstalled the application successfully."
