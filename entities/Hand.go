package entities

type Hand struct {
	Cards []Card
}

func (hand *Hand) removeCard(card Card) {
	for i, c := range hand.Cards {
		if c == card {
			hand.Cards = append(hand.Cards[:i], hand.Cards[i+1:]...)
			return
		}
	}
}

func (hand *Hand) PlayCard(card Card, boardState *BoardState) {
	hand.removeCard(card)
	switch card.GetCardType() {
	// TODO make this a function of the card, like PlayCard() on the interface
	case Land:
		hand.PlayLand(card.(BasicLand), boardState)
	case Spell:
		hand.PlaySpell(card.(BasicSpell), boardState)
	default:
		// Handle unknown card type if necessary
	}

}
func (hand *Hand) PlayLand(card BasicLand, boardState *BoardState) {
	boardState.AddLand(card)
}

func (hand *Hand) PlaySpell(card BasicSpell, boardState *BoardState) {
	tappedLands := 0
	for _, land := range boardState.Lands {
		if !land.IsTapped() && card.GetManaCost() > tappedLands {
			land.Tap()
			tappedLands++
			break
		}
	}
}

func AvailablePlays(hand Hand, boardState BoardState, turnNumber int) []Card {
	available := make([]Card, 0, len(hand.Cards))
	for _, card := range hand.Cards {
		if (card.GetCardType() == Land && boardState.CanPlayLand(turnNumber)) || (card.GetCardType() == Spell && boardState.CanPlaySpell(card)) {
			available = append(available, card)
		}
	}
	return available
}
