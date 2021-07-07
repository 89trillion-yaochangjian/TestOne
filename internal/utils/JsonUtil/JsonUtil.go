package JsonUtil

import (
	"os"
	"path/filepath"

	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Soldier struct {
	//定义士兵配置
	ID          string `json:"id"`
	Name        string `json:"Name"`
	UnlockArena string `json:"UnlockArena"`
	Rarity      string `json:"Rarity"`
	Atk         string `json:"Atk"`
	Cvc         string `json:"Cvc"`
}

var event map[string]Soldier

func ReadJsonUtil() map[string]Soldier {
	// 获取到当前目录
	pwd, _ := os.Getwd()
	//fmt.Println("当前的操作路径为:",pwd)
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "new_Soldier.json")
	//根据路径读取文件
	b, err := ioutil.ReadFile(f1)
	if err != nil {
		fmt.Print("JSON文件读取失败", err)
	}
	if err1 := json.Unmarshal(b, &event); err != nil {
		panic(err1)
	}
	return event
}

func WriteJsonUtil() {
	pwd, _ := os.Getwd() // 获取到当前目录，相当于python里的os.getcwd()
	//fmt.Println("当前的操作路径为:",pwd)
	//文件路径拼接
	f1 := filepath.Join(pwd, "conf", "new_Soldier.json")
	// 创建文件
	filePtr, err := os.Create(f1)
	// 文件创建异常处理
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	//结束关闭
	defer filePtr.Close()
	//格式化JSON数据
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		fmt.Println("格式化失败", err)
	} else {
		fmt.Println("格式化成功")
	}
	//写入文件
	filePtr.Write(data)
}
