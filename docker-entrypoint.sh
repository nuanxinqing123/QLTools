#!/bin/bash
set -e

if [ ! -s /QLTools/config/config.yaml ]; then
  echo "检测到config配置目录下不存在config.yaml，从示例文件复制一份用于初始化...\n"
  cp -fv /QLTools/sample/config.yaml /QLTools/config/config.yaml
fi

if [ -s /QLTools/config/config.yaml ]; then
  echo "检测到config配置目录下存在config.yaml，即将启动...\n"

  ./QLTools-linux-${TARGET_ARCH}

fi

exec "$@"
