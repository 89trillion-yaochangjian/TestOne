package IniUtil

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

func IniUtil() string {
	// 获取到当前目录
	pwd, _ := os.Getwd()
	fmt.Println("当前的操作路径为:", pwd)
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "app.ini")
	//解析ini文件
	cfg, err := ini.Load(f1)
	if err != nil {
		fmt.Println("文件读取失败: ", err)
	}
	//获取制定端口
	httpPort := cfg.Section("server").Key("HttpPort").String()
	if err != nil {
		fmt.Println("端口获取失败: ", err)
	}
	//返回配置文件中端口
	return httpPort
}
