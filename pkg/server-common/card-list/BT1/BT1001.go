package cardList

import "github.com/CrimsonSarah/cto/pkg/server-common"

func NewBT1001() *server.Card {
	bt1001 := &server.Card{
		ID:     "",
		Code:   "BT-001",
		Type:   "Digitama",
		Name:   "Yokomon",
		Color:  []string{"red"},
		Tribes: []string{"In-Training", "Bulb"},

		Level: 2,

		Effects: map[string]func(*server.Game){},
	}

	bt1001.Effects["WhenAttacking"] = func(*server.Game) {
		bt1001.EvolutionParent.DP += 1000
	}

	bt1001.Effects["TurnEnd"] = func(*server.Game) {
		bt1001.EvolutionParent.DP -= 1000
	}

	return bt1001
}
