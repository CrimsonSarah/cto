package turn

import (
	game "github.com/CrimsonSarah/cto/pkg/server-common"
)

const (
	Null byte = iota
	Untap
	Draw
	Breed
	Main
)

func SetUntap(game *game.Game) {
	game.TurnStep = Untap
}

func SetDraw(game *game.Game) {
	game.TurnStep = Draw
}

func SetBreed(game *game.Game) {
	game.TurnStep = Breed
}

func SetMain(game *game.Game) {
	game.TurnStep = Main
}
