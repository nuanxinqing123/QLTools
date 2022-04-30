#!/bin/sh
if [[ $(uname -a | grep "x86_64") != "" ]]
then ./QLTools-linux-amd64
else ./QLTools-linux-arm64
fi
