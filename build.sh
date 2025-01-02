#!/bin/bash

export GOOS="linux"
export GOMIPS="softfloat"

export GOARCH="mipsle"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_mipsle
echo "mipsle 平台打包完成"

export GOARCH="amd64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_amd64
echo "amd64 平台打包完成"

export GOARCH="arm"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_arm
echo "arm 平台打包完成"

export GOARCH="arm64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_arm64
echo "arm64 平台打包完成"

echo "可执行文件存放于exec目录下"