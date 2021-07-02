package soldierRouter

import (
	"AnalyseFile/internal/ctrl/soldierCrtrl"
	"AnalyseFile/internal/utils/IniUtil"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.GET("/SoldierFindALLCycUnlock", soldierCrtrl.SoldierFindALLCycUnlock)
	r.GET("/RarityFindByID", soldierCrtrl.RarityFindByID)
	r.GET("/AtkFindByID", soldierCrtrl.AtkFindByID)
	r.GET("/SoldierFindByCyc", soldierCrtrl.SoldierFindByCyc)
	r.GET("/SoldierEachStage", soldierCrtrl.SoldierEachStage)
	httpPort := IniUtil.IniUtil()
	r.Run(":" + httpPort)

}
