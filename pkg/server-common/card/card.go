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
	IsInherited     bool
}

type Digimon struct {
	Card
	IsInherited     bool
	IsTapped        bool
	InheritedEffect func()
	Effect          func()
	Security        func()
	MemoryCost      byte
	Level           byte
	DP              int
	EvolutionLine   []CardType
}

type Tamer struct {
	Card
	IsInherited     bool
	IsTapped        bool
	InheritedEffect func()
	Effect          func()
	Security        func()
	MemoryCost      byte
	EvolutionLine   []CardType
}

type Option struct {
	Card
	Effect     func()
	Security   func()
	MemoryCost byte
}

type CardType interface {
	ReturnTrigger(byte)
	Untap()
	Tap()
}

func (dt Digitama) ReturnTrigger(trigger byte) {
	dt.InheritedEffect()
}

func (dg Digimon) ReturnTrigger(trigger byte) {
	const (
		effect byte = iota
		inheritedEffect
		security
	)

	switch trigger {
	case effect:
		dg.Effect()
	case inheritedEffect:
		dg.InheritedEffect()
	case security:
		dg.Security()
	default:
	}
}

func (t Tamer) ReturnTrigger(trigger byte) {
	const (
		effect byte = iota
		inheritedEffect
		security
	)

	switch trigger {
	case effect:
		t.Effect()
	case inheritedEffect:
		t.InheritedEffect()
	case security:
		t.Security()
	default:
	}
}

func (o Option) ReturnTrigger(trigger byte) {
	const (
		effect byte = iota
		inheritedEffect
		security
	)

	switch trigger {
	case effect:
		o.Effect()
	case security:
		o.Security()
	default:
	}
}

func (dg Digimon) Untap() {
	dg.IsTapped = false
}

func (t Tamer) Untap() {
	t.IsTapped = false
}

func (dg Digimon) Tap() {
	dg.IsTapped = true
}

func (t Tamer) Tap() {
	t.IsTapped = true
}

func (dt Digitama) Tap()   {}
func (o Option) Tap()      {}
func (dt Digitama) Untap() {}
func (o Option) Untap()    {}
