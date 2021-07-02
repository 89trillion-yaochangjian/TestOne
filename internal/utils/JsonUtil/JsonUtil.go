package JsonUtil

import (
	"os"

	"encoding/json"
	"fmt"
	"io/ioutil"
)

var path = "/Users/yaochangjian/go/src/AnalyseFile/conf/new_Soldier.json"

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

func ReadJsonUtil() (map[string]Soldier) {
	//根据路径读取文件
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print("JSON文件读取失败", err)
	}
	if err1 := json.Unmarshal(b, &event); err != nil {
		panic(err1)
	}
	return event
}

func WriteJsonUtil() {
	// 创建文件
	filePtr, err := os.Create(path)
	// 文件创建异常处理
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	//结束关闭
	defer filePtr.Close()
	//格式化JSON数据
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		fmt.Println("格式化失败", err.Error())
	} else {
		fmt.Println("格式化成功")
	}
	//写入文件
	filePtr.Write(data)
}
