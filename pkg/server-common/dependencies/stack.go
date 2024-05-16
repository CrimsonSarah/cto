package dependencies

type Stack struct {
	Triggers []func()
}

func Trigger(card *Card, stack *Stack, triggerType string) {
	stack.Triggers = append(stack.Triggers, card.Effects[triggerType])
}

func Resolve(stack *Stack) {
	stack.Triggers[len(stack.Triggers)-1]()
	stack.Triggers = stack.Triggers[:len(stack.Triggers)-1]
}
