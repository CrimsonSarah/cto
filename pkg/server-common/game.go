package server

type Game struct {
	Players       [2]*Player
	TurnOwner     *Player
	TurnStep      byte
	CurrentAction byte
	Stack         *Stack
}

// Consts to facilitate code readability
const (
	Player1 byte = iota
	Player2
)
const (
	UntapStep byte = iota
	DrawStep
	BreedStep
	MainStep
	EndStep
)
const (
	Null byte = iota
	Selecting
	Targeting
	Resolving
)

// Game state related methods
func SetUntapStep(game *Game) {
	game.TurnStep = UntapStep
	for i := 0; i < len(game.TurnOwner.Board); i++ {
		Untap(game.TurnOwner.Board[i])
	}
	SetDrawStep(game)
}
func SetDrawStep(game *Game) {
	game.TurnStep = DrawStep
	Draw(game.TurnOwner, 1)
	SetBreedStep(game)
}
func SetBreedStep(game *Game) {
	game.TurnStep = BreedStep
}
func SetMainStep(game *Game) {
	game.TurnStep = MainStep
}
func SetEndStep(game *Game) {
	game.TurnStep = EndStep
}
func SetFreeAction(game *Game) {
	game.CurrentAction = Null
}
func SetSelectingAction(game *Game) {
	game.CurrentAction = Selecting
}
func SetTargetingAction(game *Game) {
	game.CurrentAction = Targeting
}
func SetResolvingAction(game *Game) {
	game.CurrentAction = Resolving
}
func ToggleTurnOwner(game *Game) {
	if game.TurnOwner == game.Players[Player1] {
		game.TurnOwner = game.Players[Player2]
	} else {
		game.TurnOwner = game.Players[Player1]
	}
}
func ResetMemory(game *Game) {
	if game.TurnOwner == game.Players[Player1] {
		game.Players[Player1].Memory = 3
		game.Players[Player2].Memory = 0
	} else {
		game.Players[Player2].Memory = 3
		game.Players[Player1].Memory = 0
	}
}
