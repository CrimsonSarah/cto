package digimon_list

import (
	"fmt"

	"github.com/CrimsonSarah/cto/pkg/server-common/card"
)

var Agumon = card.Digimon{
	Card: card.Card{
		ID:    "T-002",
		Name:  "Agumon",
		Color: []string{"red"},
	},
	Effect: func() {
		fmt.Println("Sou um agumon")
	},
	InheritedEffect: func() {
		fmt.Println("Evolui de um agumon")
	},
	IsInherited: false,
	DP:          3000,
}

var Greymon = card.Digimon{
	Card: card.Card{
		ID:    "T-003",
		Name:  "Greymon",
		Color: []string{"red"},
	},
	Effect: func() {
		fmt.Println("Sou um greymon")
	},
	InheritedEffect: func() {
		fmt.Println("Evolui de um greymon")
	},
	IsInherited: false,
	DP:          7000,
}

var MetalGreymon = card.Digimon{
	Card: card.Card{
		ID:    "T-004",
		Name:  "MetalGreymon",
		Color: []string{"red"},
	},
	Effect: func() {
		fmt.Println("Sou um metalgreymon")
	},
	InheritedEffect: func() {
		fmt.Println("Evolui de um metalgreymon")
	},
	IsInherited: false,
	DP:          11000,
}
