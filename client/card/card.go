package card

import (
	"github.com/CrimsonSarah/cto/client/digimath"
)

type Card struct {
	Name string
	Id   string
}

// A card somewhere in space
type PlacedCard struct {
	Card
	Transform digimath.Matrix44
}

func MakePlacedCard(id, name string) PlacedCard {
	return PlacedCard{
		Card: Card{
			Name: name,
			Id:   id,
		},
		Transform: digimath.Matrix44Id(),
	}
}
