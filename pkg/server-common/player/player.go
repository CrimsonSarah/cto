package player

import (
	"github.com/CrimsonSarah/cto/pkg/server-common/card"
)

type Player struct {
	ID           string
	Deck         []card.CardType
	Hand         []card.CardType
	Trash        []card.CardType
	Board        []card.CardType
	Security     []card.CardType
	Digitamas    []card.CardType
	BreedingArea []card.CardType
}

func Draw(player *Player, qt int) {
	for i := 0; i < qt; i++ {
		if len(player.Deck) > 0 {
			player.Hand = append(player.Hand, player.Deck[0])
			player.Deck = player.Deck[1:]
		}
	}
}

func Hatch(player *Player) {
	if len(player.BreedingArea) < 1 && len(player.Digitamas) > 0 {
		player.BreedingArea = append(player.BreedingArea, player.Digitamas[0])
		player.Deck = player.Digitamas[1:]
	}
}

func RemoveFromDeck(card *card.CardType, player *Player) {
}
