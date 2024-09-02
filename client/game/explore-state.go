package game

import (
	"math"

	"github.com/CrimsonSarah/cto/client/digimath"
	"github.com/CrimsonSarah/cto/client/events"
	"github.com/CrimsonSarah/cto/client/game/world"
	"github.com/CrimsonSarah/cto/client/ui"
	"github.com/gotk3/gotk3/gdk"
)

type GameExploreState struct {
	Game         *Game
	StoredCamera world.Transform

	InputForward  bool
	InputBackward bool
	InputUp       bool
	InputDown     bool
	InputLeft     bool
	InputRight    bool
}

func (s *GameExploreState) Init(game *Game) {
	s.Game = game
	s.StoredCamera = game.World.Camera
}

func (s *GameExploreState) Tick(f ui.FrameContext) {
	for ev, ok := f.Events.Dequeue(); ok; ev, ok = f.Events.Dequeue() {
		if event, ok := ev.(events.KeyDownEvent); ok {
			if event.Key == gdk.KEY_F12 {
				s.Game.ToState(&GameDefaultState{})
				return
			}

			if event.Key == gdk.KEY_W {
				s.InputForward = true
			} else if event.Key == gdk.KEY_S {
				s.InputBackward = true
			} else if event.Key == gdk.KEY_A {
				s.InputLeft = true
			} else if event.Key == gdk.KEY_D {
				s.InputRight = true
			} else if event.Key == gdk.KEY_space {
				s.InputUp = true
			} else if event.Key == gdk.KEY_Shift_L {
				s.InputDown = true
			}
		} else if event, ok := ev.(events.KeyUpEvent); ok {
			if event.Key == gdk.KEY_W {
				s.InputForward = false
			} else if event.Key == gdk.KEY_S {
				s.InputBackward = false
			} else if event.Key == gdk.KEY_A {
				s.InputLeft = false
			} else if event.Key == gdk.KEY_D {
				s.InputRight = false
			} else if event.Key == gdk.KEY_space {
				s.InputUp = false
			} else if event.Key == gdk.KEY_Shift_L {
				s.InputDown = false
			}
		} else if event, ok := ev.(events.PointerMotionEvent); ok {
			coords := s.Game.normalizedWindowCoordinates(event.X, event.Y)

			// Movement along the X axis corresponds to rotation around the
			// Y axis and movement along the Y axis corresponds to rotation
			// around the X axis.
			ry := -coords.X() * 2 * math.Pi
			rx := coords.Y() * 2 * math.Pi

			s.Game.World.Camera.Rotation = digimath.MakeVec3(rx, ry, 0)
		}
	}

	speed := float32(2)

	forward := s.Game.World.CameraForwardFlat()
	left := digimath.Matrix33RotateY(math.Pi / 2).MulV(forward)
	up := digimath.MakeVec3(0, 1, 0)
	direction := digimath.Vec3Zero

	if s.InputForward {
		direction = direction.Add(forward)
	}
	if s.InputBackward {
		direction = direction.Add(forward.Scale(-1))
	}
	if s.InputLeft {
		direction = direction.Add(left)
	}
	if s.InputRight {
		direction = direction.Add(left.Scale(-1))
	}
	if s.InputUp {
		direction = direction.Add(up)
	}
	if s.InputDown {
		direction = direction.Add(up.Scale(-1))
	}

	s.Game.World.Camera.Translate(direction.Scale(speed * f.Dtf))
}

func (s *GameExploreState) Destroy() {
	s.Game.World.Camera = s.StoredCamera
}
