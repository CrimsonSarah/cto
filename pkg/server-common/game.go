package game

import (
	"github.com/CrimsonSarah/cto/pkg/server-common/player"
	"github.com/CrimsonSarah/cto/pkg/server-common/stack"
)

type Game struct {
	Players       [2]player.Player
	TurnOwner     string
	TurnStep      byte
	CurrentAction byte
	Memory        int
	Stack         stack.Stack
}
