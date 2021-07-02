package soldierService

import "testing"

func TestAtkFindByIDHandler(t *testing.T) {
	AtkFindByIDHandlerRes := AtkFindByIDHandler("15109")
	t.Log(AtkFindByIDHandlerRes)
}
func TestSoldierEachStageHandler(t *testing.T) {
	SoldierEachStageHandlerRes := SoldierEachStageHandler("4")
	t.Log(SoldierEachStageHandlerRes)
}
func TestRarityFindByIDHandler(t *testing.T) {
	RarityFindByIDHandlerRes := RarityFindByIDHandler("15109")
	t.Log(RarityFindByIDHandlerRes)
}
func TestSoldierFindALLCycUnlockHandler(t *testing.T) {
	SoldierFindALLCycUnlockHandlerRes := SoldierFindALLCycUnlockHandler("3", "", "4")
	t.Log(SoldierFindALLCycUnlockHandlerRes)
}
func TestSoldierFindByCycHandler(t *testing.T) {
	SoldierFindByCycHandlerRes := SoldierFindByCycHandler("1000")
	t.Log(SoldierFindByCycHandlerRes)
}
