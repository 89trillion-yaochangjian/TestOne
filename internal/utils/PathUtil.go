package utils

import (
	"AnalyseFile/StructInfo"
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
)

func PathUtil() (string, *StructInfo.Response) {
	// 获取到当前目录
	pwd, err1 := os.Getwd()
	if err1 != nil {
		return "", StructInfo.FilePathErr
	}
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "new_Soldier.json")
	//获取命令行参数传入的文件路径，返回路径字符串
	var path string
	//默认读取解析后文件目录
	pflag.StringVar(&path, "path", f1, "地址")
	pflag.Parse()
	return path, nil
}
