package main

import (
	"AnalyseFile/internal/router"
	"AnalyseFile/internal/utils/JsonUtil"
	"AnalyseFile/internal/utils/PathUtil"
)

func main() {
	JsonUtil.ReadJsonUtil()
	JsonUtil.WriteJsonUtil()
	PathUtil.PathUtil()
	router.Router()
}
