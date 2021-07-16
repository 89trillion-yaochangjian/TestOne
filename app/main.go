package main

import (
	"AnalyseFile/internal/router"
	"AnalyseFile/internal/utils"
)

func main() {
	//解析文件
	utils.ReadJsonUtil()
	utils.WriteJsonUtil()
	//启动服务
	router.Router()
}
