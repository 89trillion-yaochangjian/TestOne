package router

import (
	"AnalyseFile/internal/ctrl"
	"AnalyseFile/internal/utils/IniUtil"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	//根据稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
	r.GET("/SoldierFindALLCycUnlock", ctrl.SoldierFindALLCycUnlock)
	//根据id获取稀有度
	r.GET("/RarityFindByID", ctrl.RarityFindByID)
	//根据id获取战力
	r.GET("/AtkFindByID", ctrl.AtkFindByID)
	//根据cvc获取所有合法的士兵
	r.GET("/SoldierFindByCyc", ctrl.SoldierFindByCyc)
	//根据每个阶段获取相应士兵
	r.GET("/SoldierEachStage", ctrl.SoldierEachStage)
	httpPort := IniUtil.IniUtil()
	r.Run(":" + httpPort)

}
