package service

import "testing"

func TestAtkFindByIDHandler(t *testing.T) {
	AtkFindByIDHandlerRes, err := AtkFindByIDHandler("15109")
	t.Log(AtkFindByIDHandlerRes, err)
}
func TestSoldierEachStageHandler(t *testing.T) {
	SoldierEachStageHandlerRes, err := SoldierEachStageHandler("4")
	t.Log(SoldierEachStageHandlerRes, err)
}
func TestRarityFindByIDHandler(t *testing.T) {
	RarityFindByIDHandlerRes, err := RarityFindByIDHandler("15109")
	t.Log(RarityFindByIDHandlerRes, err)
}
func TestSoldierFindALLCycUnlockHandler(t *testing.T) {
	SoldierFindALLCycUnlockHandlerRes, err := SoldierFindALLCycUnlockHandler("3", "", "4")
	t.Log(SoldierFindALLCycUnlockHandlerRes, err)
}
func TestSoldierFindByCycHandler(t *testing.T) {
	SoldierFindByCycHandlerRes, err := SoldierFindByCycHandler("1000")
	t.Log(SoldierFindByCycHandlerRes, err)
}
