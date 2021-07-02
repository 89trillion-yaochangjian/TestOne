package soldierHandler

import (
	"AnalyseFile/internal/service/soldierService"
	"AnalyseFile/internal/utils/JsonUtil"
)

/*
	根据id获取稀有度
*/

func RarityFindByIDHandler(id string) string {
	soldierMap := soldierService.GetSoldierInfo()
	var Rarity string
	for k, v := range soldierMap {
		if id == k {
			Rarity = v.Rarity
		}
	}
	return Rarity
}

/*
	根据id获取战力
*/

func AtkFindByIDHandler(id string) string {
	soldierMap := soldierService.GetSoldierInfo()
	var Atk string
	for k, v := range soldierMap {
		if id == k {
			Atk = v.Atk
		}
	}
	return Atk
}

/*
	根据稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
*/

func SoldierFindALLCycUnlockHandler(rarity string, cvc string, unlockArena string) []JsonUtil.Soldier {
	soldierMap := soldierService.GetSoldierInfo()
	var soldiers []JsonUtil.Soldier
	for _, soldier := range soldierMap {
		if rarity == soldier.Rarity && cvc == soldier.Cvc && unlockArena == soldier.UnlockArena {
			soldiers = append(soldiers, soldier)
		}
	}
	return soldiers
}

/*
	根据cvc获取所有合法的士兵
*/

func SoldierFindByCycHandler(cvc string) []JsonUtil.Soldier {
	soldierMap := soldierService.GetSoldierInfo()
	var soldiers []JsonUtil.Soldier
	for _, soldier := range soldierMap {
		if cvc == soldier.Cvc {
			soldiers = append(soldiers, soldier)
		}
	}
	return soldiers
}

/*
	根据每个阶段获取相应士兵
*/

func SoldierEachStageHandler(unlockArena string) []JsonUtil.Soldier {
	soldierMap := soldierService.GetSoldierInfo()
	var soldiers []JsonUtil.Soldier
	for _, soldier := range soldierMap {
		if unlockArena == soldier.UnlockArena {
			soldiers = append(soldiers, soldier)
		}
	}
	return soldiers
}