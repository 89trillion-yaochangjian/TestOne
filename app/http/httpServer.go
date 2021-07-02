package main

import (
	"AnalyseFile/internal/router/soldierRouter"
	"AnalyseFile/internal/utils/JsonUtil"
	"AnalyseFile/internal/utils/PathUtil"
)


func main() {
	JsonUtil.ReadJsonUtil()
	JsonUtil.WriteJsonUtil()
	PathUtil.PathUtil()
	soldierRouter.Router()
}
