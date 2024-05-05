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
	End
)

const (
	Selecting byte = iota + 6
	Targeting
	Resolving
)

func SetUntapStep(game *game.Game) {
	game.TurnStep = Untap
}

func SetDrawStep(game *game.Game) {
	game.TurnStep = Draw
}

func SetBreedStep(game *game.Game) {
	game.TurnStep = Breed
}

func SetMainStep(game *game.Game) {
	game.TurnStep = Main
}

func SetEndStep(game *game.Game) {
	game.TurnStep = End
}

func SetFreeAction(game *game.Game) {
	game.CurrentAction = Null
}

func SetSelectingAction(game *game.Game) {
	game.CurrentAction = Selecting
}

func SetTargetingAction(game *game.Game) {
	game.CurrentAction = Targeting
}

func SetResolvingAction(game *game.Game) {
	game.CurrentAction = Resolving
}

func ToggleTurnOwner(game *game.Game) {
	if game.TurnOwner == game.Players[0] {
		game.TurnOwner = game.Players[1]
	} else {
		game.TurnOwner = game.Players[0]
	}
}
