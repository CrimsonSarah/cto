package game

import (
	"github.com/CrimsonSarah/cto/pkg/server-common/card"
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

func PlayFromHand(c card.CardType, p *player.Player, game *Game) {
	stack.Trigger(c, &game.Stack, 0)
	game.Memory += c.ReturnMemoryCost()
	player.MoveFromArea(&c, 1, 3, p)
}

func PlayFromSecurity(c card.CardType, p *player.Player, game *Game) {
	stack.Trigger(c, &game.Stack, 0)
	player.MoveFromArea(&c, 5, 3, p)
}
