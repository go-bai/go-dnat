#!/bin/bash

APP_NAME=dnat
INSTALL_PATH=/usr/local/bin
DOWNLOAD_URL=https://github.com/go-bai/go-dnat/releases/download

if [[ -z $APP_VERSION  ]]; then
    echo "APP_VERSION is not set"
    exit 1
fi

arch=$(uname -m)

if [ "$arch" = "x86_64" ]; then
  curl -L -o ${APP_NAME} ${DOWNLOAD_URL}/${APP_VERSION}/dnat-linux-amd64 || { echo "Download failed"; exit 1; }
elif [ "$arch" = "aarch64" ]; then
  curl -L -o ${APP_NAME} ${DOWNLOAD_URL}/${APP_VERSION}/dnat-linux-arm64 || { echo "Download failed"; exit 1; }
else
  echo "Unknown architecture: $arch"
  exit 1
fi

chmod +x $APP_NAME
mv $APP_NAME $INSTALL_PATH