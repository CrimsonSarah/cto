package turn

import (
	game "github.com/CrimsonSarah/cto/pkg/server-common"
	card "github.com/CrimsonSarah/cto/pkg/server-common/card"
	"github.com/CrimsonSarah/cto/pkg/server-common/player"
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

const (
	Player1 byte = iota
	Player2
)

var turnowner byte

func SetUntapStep(game *game.Game) {
	game.TurnStep = Untap
	for i := 0; i < len(game.Players[turnowner].Board); i++ {
		card.CardType.Untap(game.Players[turnowner].Board[i])
	}
	SetDrawStep(game)
}

func SetDrawStep(game *game.Game) {
	game.TurnStep = Draw
	player.Draw(&game.Players[turnowner], 1)
	SetBreedStep(game)
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
	if game.TurnOwner == game.Players[Player1].ID {
		game.TurnOwner = game.Players[Player2].ID
		turnowner = Player2
	} else {
		game.TurnOwner = game.Players[Player1].ID
		turnowner = Player1
	}
}

func ResetMemory(game *game.Game) {
	if game.TurnOwner == game.Players[Player1].ID {
		game.Memory = 3
	} else {
		game.Memory = -3
	}
}
