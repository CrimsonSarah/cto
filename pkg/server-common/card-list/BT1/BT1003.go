package cardList

import "github.com/CrimsonSarah/cto/pkg/server-common"

func NewBT1003() *server.Card {
	bt1003 := &server.Card{
		ID:     "",
		Code:   "BT-003",
		Type:   "Digitama",
		Name:   "Upamon",
		Color:  []string{"blue"},
		Tribes: []string{"In-Training", "Amphibian"},

		Level: 2,

		Effects: map[string]func(*server.Game){},
	}

	bt1003.Effects["WhenAttacking"] = func(game *server.Game) {
		var opponent *server.Player
		var flag bool = false

		if game.Players[0] == bt1003.Owner {
			opponent = game.Players[1]
		} else {
			opponent = game.Players[0]
		}
		for i := 0; i < len(opponent.Board); i++ {
			if len(opponent.Board[i].EvolutionLine) == 0 {
				flag = true
			}
		}
		if flag {
			server.Draw(bt1003.Owner, 1)
		}
	}

	return bt1003
}
