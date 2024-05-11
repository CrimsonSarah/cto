package card

type Card struct {
	Code string
	Name string
}

func MakeCard(code, name string) Card {
	return Card{
		Name: name,
		Code: code,
	}
}
