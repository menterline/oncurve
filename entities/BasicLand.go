package entities

type BasicLand struct {
	isTapped bool
}

func (b BasicLand) GetManaCost() int {
	return 0
}
func (b BasicLand) GetCardType() CardType {
	return Land
}

func (b BasicLand) IsTapped() bool {
	return b.isTapped
}

func (b *BasicLand) Tap() {
	b.isTapped = true
}

func (b *BasicLand) Untap() {
	b.isTapped = false
}
