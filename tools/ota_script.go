package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func ExecRemoteScript(url, savePath string) (rst string, err error) {
	fmt.Println("拉取脚本: ", url)
	fmt.Println("脚本存放路径为: ", savePath)
	err = DownloadScript(url, savePath)
	if err != nil {
		fmt.Println("拉取脚本失败: ", err)
		return "", err
	}

	return ExecScript(savePath)
}

func DownloadScript(url, savePath string) error {
	// 拉取文件
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 建立文件用于保存数据
	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 将响应流和文件流对应起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func ExecScript(path string) (rst string, err error) {
	fmt.Println("更新脚本权限: ", path)
	err = os.Chmod(path, 7777)
	if err != nil {
		fmt.Println("更新脚本权限失败: ", err)
		return "", err
	}

	fmt.Println("执行脚本: ", path)
	cmd := exec.Command(path)
	rstByte, err := cmd.Output()
	if err != nil {
		fmt.Println("执行脚本失败: ", err)
		return "", err
	}

	return string(rstByte), nil
}
