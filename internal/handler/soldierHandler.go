package handler

import (
	"AnalyseFile/StructInfo"
	"AnalyseFile/internal/service"
)

/*
	根据id获取稀有度
*/

func RarityFindByIDHandler(id string) (string, *StructInfo.Response) {
	soldierMap, err := service.GetSoldierInfo()
	if err != nil {
		return "", err
	}
	v, ok := soldierMap[id]
	if !ok {
		return "", StructInfo.GetSoliderErr
	}
	return v.Rarity, nil
}

/*
	根据id获取战力
*/

func AtkFindByIDHandler(id string) (string, *StructInfo.Response) {
	soldierMap, err := service.GetSoldierInfo()
	if err != nil {
		return "", err
	}
	v, ok := soldierMap[id]
	if !ok {
		return "", StructInfo.GetSoliderErr
	}
	return v.Atk, nil
}

/*
	根据稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
*/

func SoldierFindALLCycUnlockHandler(rarity string, cvc string, unlockArena string) ([]StructInfo.Soldier, *StructInfo.Response) {
	soldierMap, err := service.GetSoldierInfo()
	if err != nil {
		return nil, err
	}
	var soldiers []StructInfo.Soldier
	for _, soldier := range soldierMap {
		if rarity == soldier.Rarity && cvc == soldier.Cvc && unlockArena == soldier.UnlockArena {
			soldiers = append(soldiers, soldier)
		}
	}
	return soldiers, nil
}

/*
	根据cvc获取所有合法的士兵
*/

func SoldierFindByCycHandler(cvc string) ([]StructInfo.Soldier, *StructInfo.Response) {
	soldierMap, err := service.GetSoldierInfo()
	if err != nil {
		return nil, err
	}
	var soldiers []StructInfo.Soldier
	for _, soldier := range soldierMap {
		if cvc == soldier.Cvc {
			soldiers = append(soldiers, soldier)
		}
	}
	return soldiers, nil
}

/*
	根据每个阶段获取相应士兵
*/

func SoldierEachStageHandler(unlockArena string) ([]StructInfo.Soldier, *StructInfo.Response) {
	soldierMap, err := service.GetSoldierInfo()
	if err != nil {
		return nil, err
	}
	var soldiers []StructInfo.Soldier
	for _, soldier := range soldierMap {
		if unlockArena == soldier.UnlockArena {
			soldiers = append(soldiers, soldier)
		}
	}
	return soldiers, nil
}
