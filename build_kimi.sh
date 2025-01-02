#!/bin/bash

# 构建 MIPS Little Endian 架构的可执行文件
export GOOS="linux"
export GOMIPS="softfloat"
export GOARCH="mipsle"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_mipsle
echo "MIPSLE 平台打包完成"

# 构建 AMD64 架构的可执行文件
export GOARCH="amd64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_amd64
echo "AMD64 平台打包完成"

# 构建 ARM 架构的可执行文件
export GOARCH="arm"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_arm
echo "ARM 平台打包完成"

# 构建 ARM64 架构的可执行文件
export GOARCH="arm64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_arm64
echo "ARM64 平台打包完成"

# 构建 macOS Intel 架构的可执行文件
export GOOS="darwin"
export GOARCH="amd64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_darwin_amd64
echo "macOS Intel 平台打包完成"

# 构建 macOS ARM 架构的可执行文件
export GOARCH="arm64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_darwin_arm64
echo "macOS ARM 平台打包完成"

# 构建 Windows AMD64 架构的可执行文件
export GOOS="windows"
export GOARCH="amd64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_windows_amd64.exe
echo "Windows AMD64 平台打包完成"

# 新增：构建 Windows ARM 架构的可执行文件
export GOARCH="arm"
export GOOS="windows"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_windows_arm.exe
echo "Windows ARM 平台打包完成"

echo "所有可执行文件存放于 exec 目录下"
