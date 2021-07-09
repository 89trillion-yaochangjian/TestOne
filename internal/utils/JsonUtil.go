package utils

import (
	"AnalyseFile/StructInfo"
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"os"
	"path/filepath"

	"encoding/json"
	"io/ioutil"
)

var Event map[string]StructInfo.Soldier

func ReadJsonUtil() *StructInfo.Response {
	// 获取到当前目录
	pwd, err1 := os.Getwd()
	if err1 != nil {
		log.Print("获取文件路径失败")
		return StructInfo.FilePathErr
	}
	fmt.Println(pwd)
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "newSoldier.json")
	//获取命令行参数传入的文件路径，返回路径字符串
	var path string
	//默认读取解析后文件目录
	pflag.StringVar(&path, "path", f1, "地址")
	pflag.Parse()
	//根据路径读取文件
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print("JSON文件读取失败")
		return StructInfo.FileReadErr
	}
	err2 := json.Unmarshal(b, &Event)
	if err2 != nil {
		log.Print("序列化失败")
		return StructInfo.UnmarshalErr
	}
	return nil
}

func WriteJsonUtil() *StructInfo.Response {
	// 获取到当前目录
	pwd, _ := os.Getwd()
	f1 := filepath.Join(pwd, "../conf", "newSoldier.json")
	// 创建文件
	filePtr, err := os.Create(f1)
	// 文件创建异常处理
	if err != nil {
		log.Print("文件创建失败")
		return StructInfo.FileCreateErr
	}
	//结束关闭
	defer filePtr.Close()
	//格式化JSON数据
	data, err := json.MarshalIndent(Event, "", "  ")
	if err != nil {
		log.Print("序列化失败")
		return StructInfo.UnmarshalErr
	}
	filePtr.Write(data)
	return nil
}
