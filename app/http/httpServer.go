package main

import (
	"AnalyseFile/internal/router"
	"AnalyseFile/internal/utils"
)

func main() {
	utils.ReadJsonUtil()
	utils.WriteJsonUtil()
	utils.PathUtil()
	router.Router()
}
