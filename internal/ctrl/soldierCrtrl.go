package ctrl

import (
	"AnalyseFile/internal/service"
	"AnalyseFile/internal/status"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	result "AnalyseFile/internal/model"
)

/*
	根据id获取稀有度
*/

func RarityFindByID(c *gin.Context) {
	//读取Json文件内容
	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, status.ParamErr)
		return
	}
	Rarity, err := service.RarityFindByIDHandler(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Print(err)
		return
	}
	res := result.Soldier{
		Rarity: Rarity,
	}
	c.JSON(http.StatusOK, status.OK.WithData(res))
}

/*
	根据id获取战力
*/

func AtkFindByID(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, status.ParamErr)
		return
	}
	Atk, err := service.AtkFindByIDHandler(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Print(err)
		return
	}
	res := result.Soldier{
		Atk: Atk,
	}
	c.JSON(http.StatusOK, status.OK.WithData(res))
}

/*
	根据稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
*/

func SoldierFindALLCycUnlock(c *gin.Context) {
	rarity := c.Query("rarity")
	cvc := c.Query("cvc")
	unlockArena := c.Query("unlockArena")
	if len(rarity) == 0 || len(cvc) == 0 {
		c.JSON(http.StatusInternalServerError, status.ParamErr)
		return
	}
	soldier, err := service.SoldierFindALLCycUnlockHandler(rarity, cvc, unlockArena)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, status.OK.WithData(soldier))
}

/*
	根据cvc获取所有合法的士兵
*/

func SoldierFindByCyc(c *gin.Context) {
	cvc := c.Query("cvc")
	if len(cvc) == 0 {
		c.JSON(http.StatusInternalServerError, status.ParamErr)
	}
	soldier, err := service.SoldierFindByCycHandler(cvc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, status.OK.WithData(soldier))
}

/*
	根据每个阶段获取相应士兵
*/

func SoldierEachStage(c *gin.Context) {
	unlockArena := c.Query("unlockArena")
	if len(unlockArena) == 0 {
		c.JSON(http.StatusInternalServerError, status.ParamErr)
	}
	soldier, err := service.SoldierEachStageHandler(unlockArena)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, status.OK.WithData(soldier))
}
