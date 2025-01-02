go env -w CGO_ENABLED=0 GOOS=linux GOARCH="mipsle" GOMIPS="softfloat"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_mipsle
echo "mipsle 平台打包完成"

go env -w CGO_ENABLED=0 GOOS=linux GOARCH="amd64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_amd64
echo "amd64 平台打包完成"

go env -w CGO_ENABLED=0 GOOS=linux GOARCH="arm"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_arm
echo "arm 平台打包完成"


go env -w CGO_ENABLED=0 GOOS=linux GOARCH="arm64"
go build -ldflags="-s -w" -o ./exec/mqtt_cli_arm64
echo "arm64 平台打包完成"

echo "可执行文件存放于exec目录下"