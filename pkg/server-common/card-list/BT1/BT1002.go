package cardList

import "github.com/CrimsonSarah/cto/pkg/server-common"

func NewBT1002() *server.Card {
	bt1002 := &server.Card{
		ID:     "",
		Code:   "BT-002",
		Type:   "Digitama",
		Name:   "Bebydomon",
		Color:  []string{"red"},
		Tribes: []string{"In-Training", "Baby Dragon"},

		Level: 2,

		Effects: map[string]func(*server.Game){},
	}

	bt1002.Effects["YourTurn"] = func(*server.Game) {
		for i := 0; i < len(bt1002.EvolutionParent.Tribes); i++ {
			if bt1002.EvolutionParent.Tribes[i] == "Piercing" {
				bt1002.EvolutionParent.DP += 2000
			}
		}
	}

	bt1002.Effects["YourTurnEnd"] = func(*server.Game) {
		for i := 0; i < len(bt1002.EvolutionParent.Tribes); i++ {
			if bt1002.EvolutionParent.Tribes[i] == "Piercing" {
				bt1002.EvolutionParent.DP -= 2000
			}
		}
	}

	return bt1002
}
