package cardList

import "github.com/CrimsonSarah/cto/pkg/server-common"

func NewTT001() *server.Card {
	tt001 := server.Card{
		Name:    "Arisa Kinosaki",
		Code:    "TT-001",
		Owner:   &server.Player{},
		Effects: map[string]func(*server.Game){},
	}
	tt001.Effects["StartOwnerMainPhase"] = func(game *server.Game) {
	}
	return &tt001
}
