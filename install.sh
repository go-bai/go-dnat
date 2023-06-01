#!/bin/bash

APP_NAME=dnat
INSTALL_PATH=/usr/local/bin
DOWNLOAD_URL=https://github.com/go-bai/go-dnat/releases/download
APP_VERSION=${APP_VERSION:-`curl -s GET https://api.github.com/repos/go-bai/go-dnat/tags\?per_page=1 | jq -r '.[].name'`}

if [[ -z $APP_VERSION ]]; then
  echo "APP_VERSION is not set"
  exit 1
fi

echo 1 > /proc/sys/net/ipv4/ip_forward
sed -i '/^net.ipv4.ip_forward=0/'d /etc/sysctl.conf
sed -n '/^net.ipv4.ip_forward=1/'p /etc/sysctl.conf | grep -q "net.ipv4.ip_forward=1"
if [ $? -ne 0 ]; then
  echo -e "net.ipv4.ip_forward=1" >> /etc/sysctl.conf && sysctl -p
fi

arch=$(uname -m)

if [ "$arch" = "x86_64" ]; then
  curl -L -o ${APP_NAME} ${DOWNLOAD_URL}/${APP_VERSION}/dnat-linux-amd64 || {
    echo "Download failed"
    exit 1
  }
elif [ "$arch" = "aarch64" ]; then
  curl -L -o ${APP_NAME} ${DOWNLOAD_URL}/${APP_VERSION}/dnat-linux-arm64 || {
    echo "Download failed"
    exit 1
  }
else
  echo "Unknown architecture: $arch"
  exit 1
fi

chmod +x $APP_NAME
mv $APP_NAME $INSTALL_PATH