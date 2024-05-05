package game

type Game struct {
	Players       [2]string
	TurnOwner     string
	TurnStep      byte
	CurrentAction byte
}
