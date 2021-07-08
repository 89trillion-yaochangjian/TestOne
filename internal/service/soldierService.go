package service

import (
	"AnalyseFile/StructInfo"
	"AnalyseFile/internal/utils"
)

/*
	获取士兵信息
*/

var SoldierMap map[string]StructInfo.Soldier

func GetSoldierInfo() (map[string]StructInfo.Soldier, *StructInfo.Response) {

	//读取Json文件内容
	if len(SoldierMap) == 0 {
		soldierMapFisrtRece, err := utils.ReadJsonUtil()
		SoldierMap = soldierMapFisrtRece
		return SoldierMap, err
	}
	return SoldierMap, nil
}
