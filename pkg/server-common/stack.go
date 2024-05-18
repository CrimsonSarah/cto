package server

type Stack struct {
	Triggers []func(*Game)
}

func Trigger(card *Card, game *Game, triggerType string) {
	game.Stack.Triggers = append(game.Stack.Triggers, card.Effects[triggerType])
}

func Resolve(game *Game) {
	game.Stack.Triggers[len(game.Stack.Triggers)-1](game)
	game.Stack.Triggers = game.Stack.Triggers[:len(game.Stack.Triggers)-1]
}
