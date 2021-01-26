#!/bin/bash

mkdir $HOME/.go-pass

rm -rf go-pass;

git clone https://github.com/victorlpgazolli/go-pass.git;

echo "We need root permissions to move the executable to /usr/sbin"

if ! [ $(id -u) = 0 ]; then
   echo "This script must be run as root" 
   exit 1
fi

mv go-pass/pass /usr/sbin;

rm -rf go-pass;



echo 'pass installed in /usr/sbin'