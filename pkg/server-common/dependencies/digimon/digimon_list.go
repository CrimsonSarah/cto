package digimon_list

import (
	"fmt"

	"github.com/CrimsonSarah/cto/pkg/server-common/dependencies"
)

var Agumon = dependencies.Card{

	Code:  "T-002",
	Name:  "Agumon",
	Color: []string{"red"},
	Level: 3,
	Effects: map[string]func(){
		"OnPlay": func() {
			fmt.Println("Não tenho linha evolutiva")
		},
		"WhenAttacking": func() {
			fmt.Println("Xama neeneeeeeem")
		},
	},
	IsInherited: false,
	DP:          3000,
}
