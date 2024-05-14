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
	MemoryCost      int
	InheritedEffect func()
	OnPlay          func()
	Effect          func()
	Security        func()
	Level           byte
	DP              int
	EvolutionLine   []CardType
}

type Tamer struct {
	Card
	IsInherited     bool
	IsTapped        bool
	MemoryCost      int
	InheritedEffect func()
	OnPlay          func()
	Effect          func()
	Security        func()
	EvolutionLine   []CardType
}

type Option struct {
	Card
	MemoryCost int
	OnPlay     func()
	Security   func()
}

type CardType interface {
	ReturnTrigger(byte)
	ReturnMemoryCost() int
	Untap()
	Tap()
}

func (dg Digimon) ReturnTrigger(trigger byte) {
	const (
		onplay byte = iota
		effect
		inheritedEffect
		security
	)

	switch trigger {
	case onplay:
		dg.OnPlay()
	case effect:
		dg.Effect()
	case inheritedEffect:
		dg.InheritedEffect()
	case security:
		dg.Security()
	default:
	}
}
func (dg Digimon) ReturnMemoryCost() int { return dg.MemoryCost }
func (dg Digimon) Untap()                { dg.IsTapped = false }
func (dg Digimon) Tap()                  { dg.IsTapped = true }

func (t Tamer) ReturnTrigger(trigger byte) {
	const (
		onplay byte = iota
		effect
		inheritedEffect
		security
	)

	switch trigger {
	case onplay:
		t.OnPlay()
	case effect:
		t.Effect()
	case inheritedEffect:
		t.InheritedEffect()
	case security:
		t.Security()
	default:
	}
}
func (t Tamer) ReturnMemoryCost() int { return t.MemoryCost }
func (t Tamer) Untap()                { t.IsTapped = false }
func (t Tamer) Tap()                  { t.IsTapped = true }

func (o Option) ReturnTrigger(trigger byte) {
	const (
		onplay byte = iota
		effect
		inheritedEffect
		security
	)

	switch trigger {
	case effect:
		o.OnPlay()
	case security:
		o.Security()
	default:
	}
}
func (o Option) ReturnMemoryCost() int { return o.MemoryCost }
func (o Option) Untap()                {}
func (o Option) Tap()                  {}

func (dt Digitama) ReturnTrigger(trigger byte) { dt.InheritedEffect() }
func (dt Digitama) ReturnMemoryCost() int      { return 0 }
func (dt Digitama) Untap()                     {}
func (dt Digitama) Tap()                       {}
