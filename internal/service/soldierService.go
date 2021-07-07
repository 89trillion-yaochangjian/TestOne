package service

import (
	"AnalyseFile/internal/utils/JsonUtil"
)

/*
	获取士兵信息
*/
var soldierMap map[string]JsonUtil.Soldier

func GetSoldierInfo() map[string]JsonUtil.Soldier {

	//读取Json文件内容
	if len(soldierMap) == 0 {
		soldierMap = JsonUtil.ReadJsonUtil()
	}
	return soldierMap
}
