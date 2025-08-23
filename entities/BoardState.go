package entities

type BoardState struct {
	TurnIndex int
	Lands     []BasicLand
}

func (b *BoardState) AddLand(land BasicLand) {
	b.Lands = append(b.Lands, land)
}

func (b *BoardState) NextTurn() {
	b.TurnIndex++
}

func (b BoardState) CanPlaySpell(card Card) bool {
	untappedLands := 0
	for _, land := range b.Lands {
		if !land.IsTapped() {
			untappedLands++
		}
	}
	if card.GetManaCost() <= untappedLands {
		return true
	}
	return false
}

// turnNumber is 0-indexed
func (b BoardState) CanPlayLand(turnNumber int) bool {
	if len(b.Lands) <= turnNumber {
		return true
	}
	return false
}

func (b *BoardState) UntapLands() {
	for i := range b.Lands {
		b.Lands[i].Untap()
	}
}

func (b *BoardState) GetNumUntappedLands() int {
	numUntapped := 0
	for _, land := range b.Lands {
		if !land.IsTapped() {
			numUntapped++
		}
	}
	return numUntapped
}
