package entities

type BasicSpell struct {
	cost int
}

func NewBasicSpell(cost int) BasicSpell {
	return BasicSpell{cost: cost}
}

func (b BasicSpell) GetManaCost() int {
	return b.cost
}
func (b BasicSpell) GetCardType() CardType {
	return Spell
}

func (card BasicSpell) Play(boardState *BoardState) {
	tappedLands := 0
	for _, land := range boardState.Lands {
		if !land.IsTapped() && card.GetManaCost() > tappedLands {
			land.Tap()
			tappedLands++
			break
		}
	}
}
