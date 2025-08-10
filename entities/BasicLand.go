package entities

type BasicLand struct {
}

func (b BasicLand) GetManaCost() int {
	return 0
}
func (b BasicLand) GetCardType() CardType {
	return Land
}
