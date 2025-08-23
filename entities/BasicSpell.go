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

func (card BasicSpell) CanPlay(boardState BoardState) bool {
	untappedLands := 0
	for _, land := range boardState.Lands {
		if !land.IsTapped() {
			untappedLands++
		}
	}
	if card.GetManaCost() <= untappedLands {
		return true
	}
	return false
}

func (card BasicSpell) Play(boardState *BoardState) {
	boardState.TapLands(card.GetManaCost())
}
