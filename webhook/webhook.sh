#!/bin/bash

svr="httpServer"
wordDir="webhookTmp"
blogName="hankshell"
confName="config.toml"
assertName="static_data"

echo "---webhook start---"
# 重新编译项目
git pull

#清除老进程
pkill ${svr}

#开启新进程
pwd
go build  -o ${svr} ./main/main.go 
nohup ./${svr} &

#更新submodule
git submodule update --init --recursive

# 拉去github项目
cd webhook
rm -rf Note
git clone https://github.com/Hank00AAA/Note.git
rm -rf hankshell/content/notes
pwd
cp Note/*  ../hankshell/content/notes/
cd ../hankshell
hugo
cd ../
rm -rf static_data/*
mv hankshell/public/* static_data/

echo "---webhook finish----"