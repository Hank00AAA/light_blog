#!/bin/bash

svr="httpServer"
wordDir="webhookTmp"
blogName="hankshell"
confName="config.toml"
assertName="static_data"

echo "---webhook start---"
# 
git pull

# 删除submodule
rm  ./hankshell/content/notes/*

# 重新编译
rm -f /usr/local/bin/${svr}
go build  -o /usr/local/bin/${svr} ./main/main.go

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
