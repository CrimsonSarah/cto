package digimon_list

import (
	"fmt"
	"testing"

	"github.com/CrimsonSarah/cto/pkg/server-common/dependencies"
)

var Bibimon = dependencies.Card{

	Code:  "BT6-003",
	Name:  "Bibimon",
	Color: []string{"yellow"},
	Level: 2,
	Type:  "Digitama",
	Owner: &dependencies.Player{
		ID:       "playerteste1",
		Security: []*dependencies.Card{},
	},
	Effects:     map[string]func(){},
	IsInherited: false,
}

func SetBibimonEffects() {
	// Since it isn't possible to reference Bibimon's owner
	//in a function inside Bibimon declaration,
	//Bibimon's effects are declared in a separate method
	//so it can properly read it's owners fields
	Bibimon.Effects["WhenAttacking"] = func() {
		if len(Bibimon.Owner.Security) > 2 {
			dependencies.Draw(Bibimon.Owner, 1)
		}
	}
}

func TestBibimon(t *testing.T) {
	//Set bibimon's effect so it actually do something when triggered
	SetBibimonEffects()

	//Add 2 bibimon to owners security, so it shouldn't actually trigger the draw portion yet
	Bibimon.Owner.Security = append(Bibimon.Owner.Security, &Bibimon)
	Bibimon.Owner.Security = append(Bibimon.Owner.Security, &Bibimon)

	//Add some bibimon to owners deck, so they can draw something
	Bibimon.Owner.Deck = append(Bibimon.Owner.Security, &Bibimon)
	Bibimon.Owner.Deck = append(Bibimon.Owner.Security, &Bibimon)
	Bibimon.Owner.Deck = append(Bibimon.Owner.Security, &Bibimon)

	//Declares new stack (usually a &game.Stack would be used to store triggers)
	var newstack = dependencies.Stack{}

	//Triggers and resolves bibimon when attacking effect
	dependencies.Trigger(&Bibimon, &newstack, "WhenAttacking")
	dependencies.Resolve(&newstack)

	//Check if owners hand is empty (it should be)
	fmt.Println(Bibimon.Owner.Hand)

	//Add one last bibimon to security, so the draw portion should trigger now
	Bibimon.Owner.Security = append(Bibimon.Owner.Security, &Bibimon)

	//Triggers and resolves bibimon when attaking effect
	dependencies.Trigger(&Bibimon, &newstack, "WhenAttacking")
	dependencies.Resolve(&newstack)

	//Check if there's a new card in hand (there should be)
	fmt.Println(Bibimon.Owner.Hand)
}
