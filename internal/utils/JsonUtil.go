package utils

import (
	"AnalyseFile/StructInfo"
	"os"
	"path/filepath"

	"encoding/json"
	"io/ioutil"
)

var event map[string]StructInfo.Soldier

func ReadJsonUtil() (map[string]StructInfo.Soldier, *StructInfo.Response) {
	// 获取到当前目录
	pwd, err1 := os.Getwd()
	if err1 != nil {
		return event, StructInfo.FilePathErr
	}
	//fmt.Println("当前的操作路径为:",pwd)
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "new_Soldier.json")
	//根据路径读取文件
	b, err := ioutil.ReadFile(f1)
	if err != nil {
		return event, StructInfo.FileReadErr
	}
	err2 := json.Unmarshal(b, &event)
	if err2 != nil {
		return event, StructInfo.UnmarshalErr
	}
	return event, nil
}

func WriteJsonUtil() *StructInfo.Response {
	pwd, _ := os.Getwd() // 获取到当前目录，相当于python里的os.getcwd()
	//fmt.Println("当前的操作路径为:",pwd)
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "new_Soldier.json")
	// 创建文件
	filePtr, err := os.Create(f1)
	// 文件创建异常处理
	if err != nil {
		return StructInfo.FileCreateErr
	}
	//结束关闭
	defer filePtr.Close()
	//格式化JSON数据
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return StructInfo.UnmarshalErr
	}
	//写入文件
	filePtr.Write(data)
	return nil
}
