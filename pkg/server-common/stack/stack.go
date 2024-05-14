package stack

import (
	"github.com/CrimsonSarah/cto/pkg/server-common/card"
)

type Stack struct {
	Triggers     []card.CardType
	TriggerTypes []byte
}

func Trigger(card card.CardType, stack *Stack, triggerType byte) {
	stack.Triggers = append(stack.Triggers, card)
	stack.TriggerTypes = append(stack.TriggerTypes, triggerType)
}

func Resolve(stack *Stack) {
	stack.Triggers[len(stack.Triggers)-1].ReturnTrigger(stack.TriggerTypes[len(stack.Triggers)-1])
	stack.Triggers = stack.Triggers[:len(stack.Triggers)-1]
	stack.TriggerTypes = stack.TriggerTypes[:len(stack.TriggerTypes)-1]
}
