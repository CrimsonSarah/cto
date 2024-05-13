package stack

import (
	"github.com/CrimsonSarah/cto/pkg/server-common/card"
)

type Stack struct {
	Triggers []card.CardType
}

func Trigger(card card.CardType, stack *Stack) {
	stack.Triggers = append(stack.Triggers, card)
}

func Resolve(stack *Stack, triggerType byte) {
	stack.Triggers[len(stack.Triggers)-1].ReturnTrigger(triggerType)
	stack.Triggers = stack.Triggers[:len(stack.Triggers)-1]
}
