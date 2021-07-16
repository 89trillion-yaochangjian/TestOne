package model

type Soldier struct {
	//定义士兵配置
	ID          string `json:"id"`
	Name        string `json:"Name"`
	UnlockArena string `json:"UnlockArena"`
	Rarity      string `json:"Rarity"`
	Atk         string `json:"Atk"`
	Cvc         string `json:"Cvc"`
}
