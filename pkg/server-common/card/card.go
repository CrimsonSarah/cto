package card

type Card struct {
	ID          string
	Name        string
	Color       []string
	Description string
	Keywords    []string
	Tribes      []string
}

type Digitama struct {
	Card
	InheritedEffect func()
}

type Digimon struct {
	Card
	InheritedEffect func()
	Effect          func()
	Security        func()
	MemoryCost      byte
	Level           byte
	DP              int
}

type Tamer struct {
	Card
	Effect     func()
	Security   func()
	MemoryCost byte
}

type Option struct {
	Card
	Effect     func()
	Security   func()
	MemoryCost byte
}

type CardType interface {
	ReturnTrigger(trigger byte)
}

func (dt Digitama) ReturnTrigger(trigger byte) func() {
	return dt.InheritedEffect
}

func (dg Digimon) ReturnTrigger(trigger byte) func() {
	const (
		effect byte = iota
		inheritedEffect
		security
	)

	switch trigger {
	case effect:
		return dg.Effect
	case inheritedEffect:
		return dg.InheritedEffect
	case security:
		return dg.Security
	default:
		return nil
	}
}

func (t Tamer) ReturnTrigger(trigger byte) func() {
	const (
		effect byte = iota
		security
	)

	switch trigger {
	case effect:
		return t.Effect
	case security:
		return t.Security
	default:
		return nil
	}
}
func (o Option) ReturnTrigger(trigger byte) func() {
	const (
		effect byte = iota
		security
	)

	switch trigger {
	case effect:
		return o.Effect
	case security:
		return o.Security
	default:
		return nil
	}
}
