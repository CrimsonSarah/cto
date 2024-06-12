package dependencies

type Card struct {
	//Card identifiers
	ID         string
	Code       string
	Type       string
	Name       string
	Color      []string
	Tribes     []string
	MemoryCost byte

	//State relative fields
	IsTapped        bool
	IsInherited     bool
	EvolutionLine   []*Card
	EvolutionParent *Card
	Owner           *Player

	//Digimon unique fields
	Level byte
	DP    int

	//Name keys according to effect timing (i.e.: WhenDigivolving, OnPlay, EndAttack, etc.)
	Effects map[string]func()
}
