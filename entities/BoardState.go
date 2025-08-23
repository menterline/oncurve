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
	b.untapLands()
}

func (b *BoardState) untapLands() {
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
