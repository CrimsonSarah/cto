package turn

import (
	"fmt"
	"testing"

	game "github.com/CrimsonSarah/cto/pkg/server-common"
	"github.com/stretchr/testify/assert"
)

func TestTurn(t *testing.T) {
	joguinhoteste := game.Game{
		Players:       [2]string{"idteste1", "idteste2"},
		TurnOwner:     "idteste1",
		TurnStep:      0,
		CurrentAction: 0,
	}
	SetUntapStep(&joguinhoteste)
	SetResolvingAction(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Untap)
	fmt.Print(joguinhoteste.TurnOwner)
	fmt.Print(joguinhoteste.TurnStep)
	fmt.Print(joguinhoteste.CurrentAction)

	SetDrawStep(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Draw)
	fmt.Print(joguinhoteste.TurnStep)
	fmt.Print(joguinhoteste.CurrentAction)

	SetBreedStep(&joguinhoteste)
	SetFreeAction(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Breed)
	fmt.Print(joguinhoteste.TurnStep)
	fmt.Print(joguinhoteste.CurrentAction)

	SetMainStep(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Main)
	fmt.Print(joguinhoteste.TurnStep)
	fmt.Print(joguinhoteste.CurrentAction)

	SetEndStep(&joguinhoteste)
	SetResolvingAction(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, End)
	fmt.Print(joguinhoteste.TurnStep)
	fmt.Print(joguinhoteste.CurrentAction)

	ToggleTurnOwner(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnOwner, joguinhoteste.Players[1])

	SetUntapStep(&joguinhoteste)
	SetResolvingAction(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Untap)
	fmt.Print(joguinhoteste.TurnOwner)
	fmt.Print(joguinhoteste.TurnStep)
	fmt.Print(joguinhoteste.CurrentAction)
}
