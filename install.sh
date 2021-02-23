#!/usr/bin/env bash

wget https://github.com/stephenharris/s3album/releases/download/0.1.0/s3album-0.1.0-linux-arm.tar.gz -O s3album.tar.gz
tar -xvf s3album.tar.gz
mv s3album /usr/local/bin/s3album
rm -rf s3album.tar.gz

mkdir -p /home/pi/.config/s3album
touch /home/pi/.config/s3album/s3album.yaml
echo """bucket: 
  name: $1
  region: $2""" > /home/pi/.config/s3album/s3album.yaml


touch /lib/systemd/system/s3album.service

echo """[Unit]
Description=s3album

[Service]
Type=simple
User=pi
Restart=always
RestartSec=5s
ExecStart=/usr/local/bin/s3album

[Install]
WantedBy=multi-user.target""" > /lib/systemd/system/s3album.service

service s3album start
systemctl enable s3album.service