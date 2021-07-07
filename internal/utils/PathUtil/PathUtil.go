package PathUtil

import (
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
)

func PathUtil() string {
	// 获取到当前目录
	pwd, _ := os.Getwd()
	//fmt.Println("当前的操作路径为:",pwd)
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "new_Soldier.json")
	//获取命令行参数传入的文件路径，返回路径字符串
	var path string
	//默认读取解析后文件目录
	pflag.StringVar(&path, "path", f1, "地址")
	pflag.Parse()
	return path
}
