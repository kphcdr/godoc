#!/bin/bash

echo "拉取最新代码中。。。"

git pull

echo "打包程序。。。"

go build

echo "kill 应用。。。"

ps -ef |grep "showdoc" | grep -v 'grep' | awk '{print $2}' | xargs kill -9

echo "启动应用。。。"

nohup /root/go/src/showdoc/showdoc &

