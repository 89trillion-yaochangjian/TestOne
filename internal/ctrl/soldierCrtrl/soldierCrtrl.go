package soldierCrtrl

import (
	"AnalyseFile/internal/handler/soldierHandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	根据id获取稀有度
*/

func RarityFindByID(c *gin.Context) {
	//读取Json文件内容
	id := c.Query("id")
	Rarity := soldierHandler.RarityFindByIDHandler(id)
	c.JSON(http.StatusOK, gin.H{
		"Rarity": Rarity,
	})
}

/*
	根据id获取战力
*/

func AtkFindByID(c *gin.Context) {
	id := c.Query("id")
	Atk := soldierHandler.AtkFindByIDHandler(id)
	c.JSON(http.StatusOK, gin.H{
		"Atk": Atk,
	})
}

/*
	根据稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
*/

func SoldierFindALLCycUnlock(c *gin.Context) {
	rarity := c.Query("rarity")
	cvc := c.Query("cvc")
	unlockArena := c.Query("unlockArena")
	soldier := soldierHandler.SoldierFindALLCycUnlockHandler(rarity, cvc, unlockArena)
	c.JSON(http.StatusOK,soldier)
}

/*
	根据cvc获取所有合法的士兵
*/

func SoldierFindByCyc(c *gin.Context) {
	cvc := c.Query("cvc")
	soldier := soldierHandler.SoldierFindByCycHandler(cvc)
	c.JSON(http.StatusOK, soldier)
}

/*
	根据每个阶段获取相应士兵
*/

func SoldierEachStage(c *gin.Context) {
	unlockArena := c.Query("unlockArena")
	soldier := soldierHandler.SoldierEachStageHandler(unlockArena)
	c.JSON(http.StatusOK, soldier)

}

