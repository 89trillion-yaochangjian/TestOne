package IniUtil

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func IniUtil() string {
	//解析ini文件
	cfg, err := ini.Load("/Users/yaochangjian/go/src/AnalyseFile/conf/app.ini")
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
