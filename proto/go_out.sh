#!/bin/sh

echo "start"

proto_input=""
proto_out=""

for file in proto/*; do
    # proto文件,如proto/base.proto
    proto_input+=" $file"

    # 从左向右截取第一个string后的字符串
    pb_file=${file#*proto/}

    # 使用$replacement, 来代替第一个匹配的$substring
    pb_file=${pb_file/proto/pb.go}

    # 输出的文件
    proto_out+=" $pb_file"
done

#echo $proto_input
#echo $proto_out

# protoc 命令
# protoc --go_out=plugins=micro:.. proto/base.proto proto/account.proto
set -eux # 显示执行命令
protoc --go_out=plugins=micro:.. $proto_input
set +eux

# 添加到git
git add *

echo "done"