package game

import (
	"github.com/CrimsonSarah/cto/client/events"
	"github.com/CrimsonSarah/cto/client/ui"
	"github.com/gotk3/gotk3/gdk"
)

type GameDefaultState struct {
	Game *Game
}

func (s *GameDefaultState) Init(game *Game) {
	s.Game = game
}

func (s *GameDefaultState) Tick(f ui.FrameContext) {
	for ev, ok := f.Events.Dequeue(); ok; ev, ok = f.Events.Dequeue() {
		if event, ok := ev.(events.KeyDownEvent); ok {
			if event.Key == gdk.KEY_F12 {
				s.Game.ToState(&GameExploreState{})
			}
		}
	}
}

func (s *GameDefaultState) Destroy() {}
