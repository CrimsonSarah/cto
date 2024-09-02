package game

import "github.com/CrimsonSarah/cto/client/ui"

type GameState interface {
	Init(game *Game)
	Tick(f ui.FrameContext)
	Destroy()
}

func (game *Game) ToState(state GameState) {
	game.State.Destroy()
	state.Init(game)
	game.State = state
}
