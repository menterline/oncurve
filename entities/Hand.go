package entities

type Hand struct {
	Cards []Card
}

func (hand *Hand) removeCardFromHand(card Card) {
	for i, c := range hand.Cards {
		if c == card {
			hand.Cards = append(hand.Cards[:i], hand.Cards[i+1:]...)
			return
		}
	}
}

func (hand *Hand) PlayCard(card Card, boardState *BoardState) {
	hand.removeCardFromHand(card)
	card.Play(boardState)
}

func AvailablePlays(hand Hand, boardState BoardState, turnNumber int) []Card {
	available := make([]Card, 0, len(hand.Cards))
	for _, card := range hand.Cards {
		// TODO can add interface method on card canPlay(boardState)bool
		if (card.GetCardType() == Land && boardState.CanPlayLand(turnNumber)) || (card.GetCardType() == Spell && boardState.CanPlaySpell(card)) {
			available = append(available, card)
		}
	}
	return available
}
