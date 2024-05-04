package turn

import (
	"fmt"
	"testing"

	game "github.com/CrimsonSarah/cto/pkg/server-common"
	"github.com/stretchr/testify/assert"
)

func TestTurn(t *testing.T) {
	joguinhoteste := game.Game{
		TurnStep: 0,
	}
	SetUntap(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Untap)
	fmt.Print(joguinhoteste.TurnStep)

	SetDraw(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Draw)
	fmt.Print(joguinhoteste.TurnStep)

	SetBreed(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Breed)
	fmt.Print(joguinhoteste.TurnStep)

	SetMain(&joguinhoteste)
	assert.Equal(t, joguinhoteste.TurnStep, Main)
	fmt.Print(joguinhoteste.TurnStep)
}
