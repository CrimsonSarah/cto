package game

import (
	stack "github.com/CrimsonSarah/cto/pkg/server-common/stack"
)

type Game struct {
	Players       [2]string
	TurnOwner     string
	TurnStep      byte
	CurrentAction byte
	Stack         stack.Stack
}
