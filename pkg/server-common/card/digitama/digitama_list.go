package digitama_list

import (
	"fmt"

	"github.com/CrimsonSarah/cto/pkg/server-common/card"
)

var Koromon = card.Digitama{
	Card: card.Card{
		ID:    "T-001",
		Name:  "Koromon",
		Color: []string{"red"},
	},
	InheritedEffect: func() {
		fmt.Println("Sou digitama")
	},
	IsInherited: false,
}
