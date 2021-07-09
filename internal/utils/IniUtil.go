package utils

import (
	"AnalyseFile/StructInfo"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

func IniUtil() (string, *StructInfo.Response) {
	// 获取到当前目录
	pwd, err1 := os.Getwd()
	if err1 != nil {
		return "", StructInfo.FilePathErr
	}
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "app.ini")
	//解析ini文件
	cfg, err := ini.Load(f1)
	if err != nil {
		log.Print("JSON文件读取失败")
		return "", StructInfo.FileReadErr
	}
	//获取制定端口
	httpPort := cfg.Section("server").Key("HttpPort").String()

	//返回配置文件中端口
	return httpPort, nil
}
