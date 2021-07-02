package PathUtil

import (
	"github.com/spf13/pflag"
)

var pahtValue = "/Users/yaochangjian/go/src/AnalyseFile/conf/new_Soldier.json"

func PathUtil() string {
	//获取命令行参数传入的文件路径，返回路径字符串
	var path string
	//默认读取解析后文件目录
	pflag.StringVar(&path, "path", pahtValue, "地址")
	pflag.Parse()
	return path
}
