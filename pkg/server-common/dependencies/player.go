package dependencies

type Player struct {
	ID           string
	Deck         []*Card
	Hand         []*Card
	Trash        []*Card
	Board        []*Card
	Security     []*Card
	Digitamas    []*Card
	BreedingArea []*Card
	Memory       byte
}

// Methods for player related actions
func Draw(player *Player, qt int) {
	for i := 0; i < qt; i++ {
		if len(player.Deck) > 0 {
			player.Hand = append(player.Hand, player.Deck[0])
			player.Deck = player.Deck[1:]
		}
	}
}
func Hatch(player *Player) {
	if len(player.BreedingArea) < 1 && len(player.Digitamas) > 0 {
		player.BreedingArea = append(player.BreedingArea, player.Digitamas[0])
		player.Deck = player.Digitamas[1:]
	}
}
func MoveFromArea(card *Card, area byte, target byte, player *Player) {
	const (
		deck byte = iota
		hand
		trash
		board
		breeding
		security
		decktop
		deckbottom
	)

	switch area {

	case deck:
		switch target {

		case hand:
			for i := 0; i < len(player.Deck); i++ {
				if card == player.Deck[i] {
					player.Hand = append(player.Hand, player.Deck[i])
					player.Deck = append(player.Deck[:i], player.Deck[i+1:]...)
				}
			}

		case trash:
			for i := 0; i < len(player.Deck); i++ {
				if card == player.Deck[i] {
					player.Trash = append(player.Trash, player.Deck[i])
					player.Deck = append(player.Deck[:i], player.Deck[i+1:]...)
				}
			}

		case board:
			for i := 0; i < len(player.Deck); i++ {
				if card == player.Deck[i] {
					player.Board = append(player.Board, player.Deck[i])
					player.Deck = append(player.Deck[:i], player.Deck[i+1:]...)
				}
			}

		case security:
			for i := 0; i < len(player.Deck); i++ {
				if card == player.Deck[i] {
					player.Security = append(player.Security, player.Deck[i])
					player.Deck = append(player.Deck[:i], player.Deck[i+1:]...)
				}
			}
		}

	case hand:
		switch target {

		case deckbottom:
			for i := 0; i < len(player.Hand); i++ {
				if card == player.Hand[i] {
					player.Deck = append(player.Deck, player.Hand[i])
					player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
				}
			}

		case decktop:
			for i := 0; i < len(player.Hand); i++ {
				if card == player.Hand[i] {
					player.Deck = append(player.Hand[i:i], player.Deck[0:]...)
					player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
				}
			}

		case trash:
			for i := 0; i < len(player.Hand); i++ {
				if card == player.Hand[i] {
					player.Trash = append(player.Trash, player.Hand[i])
					player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
				}
			}

		case board:
			for i := 0; i < len(player.Hand); i++ {
				if card == player.Hand[i] {
					player.Board = append(player.Board, player.Hand[i])
					player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
				}
			}

		case security:
			for i := 0; i < len(player.Hand); i++ {
				if card == player.Hand[i] {
					player.Security = append(player.Security, player.Hand[i])
					player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
				}
			}

		case breeding:
			for i := 0; i < len(player.Hand); i++ {
				if card == player.Hand[i] {
					player.Security = append(player.Security, player.Hand[i])
					player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
				}
			}
		}

	case board:
		switch target {

		case deckbottom:
			for i := 0; i < len(player.Board); i++ {
				if card == player.Board[i] {
					player.Deck = append(player.Deck, player.Board[i])
					player.Board = append(player.Board[:i], player.Board[i+1:]...)
				}
			}

		case decktop:
			for i := 0; i < len(player.Board); i++ {
				if card == player.Board[i] {
					player.Deck = append(player.Board[i:i], player.Deck[0:]...)
					player.Board = append(player.Board[:i], player.Board[i+1:]...)
				}
			}

		case trash:
			for i := 0; i < len(player.Board); i++ {
				if card == player.Board[i] {
					player.Trash = append(player.Trash, player.Board[i])
					player.Board = append(player.Board[:i], player.Board[i+1:]...)
				}
			}

		case hand:
			for i := 0; i < len(player.Board); i++ {
				if card == player.Board[i] {
					player.Hand = append(player.Hand, player.Board[i])
					player.Board = append(player.Board[:i], player.Board[i+1:]...)
				}
			}

		case security:
			for i := 0; i < len(player.Board); i++ {
				if card == player.Board[i] {
					player.Security = append(player.Security, player.Board[i])
					player.Board = append(player.Board[:i], player.Board[i+1:]...)
				}
			}
		}

	case breeding:
		player.BreedingArea = append(player.Board, player.BreedingArea[0])
		player.BreedingArea = player.BreedingArea[1:]

	case security:
		switch target {

		case trash:
			for i := 0; i < len(player.Security); i++ {
				if card == player.Security[i] {
					player.Trash = append(player.Trash, player.Security[i])
					player.Security = append(player.Security[:i], player.Security[i+1:]...)
				}
			}

		case hand:
			for i := 0; i < len(player.Security); i++ {
				if card == player.Security[i] {
					player.Hand = append(player.Hand, player.Security[i])
					player.Security = append(player.Security[:i], player.Security[i+1:]...)
				}
			}

		case board:
			for i := 0; i < len(player.Security); i++ {
				if card == player.Security[i] {
					player.Board = append(player.Board, player.Security[i])
					player.Security = append(player.Security[:i], player.Security[i+1:]...)
				}
			}
		}
	}
}
func Recover(player *Player) {
	player.Security = append(player.Security, player.Deck[0])
	player.Deck = player.Deck[1:]
}
func PlayFromHand(card *Card, player *Player, stack *Stack) {
	Trigger(card, stack, "OnPlay")
	player.Memory -= card.MemoryCost
	MoveFromArea(card, 1, 3, player)
}
func PlayFromSecurity(card *Card, player *Player, stack *Stack) {
	Trigger(card, stack, "OnPlay")
	MoveFromArea(card, 5, 3, player)
}
func Tap(card *Card)   { card.IsTapped = true }
func Untap(card *Card) { card.IsTapped = false }
