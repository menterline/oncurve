package entities

import "errors"

type Deck struct {
	cards []Card
}

/*
when building the deck, we need a list of spells, and a list of lands,
then randomly merge them together
*/
func NewDeck(s SimSettingsData) Deck {
	return Deck{cards: make([]Card, 60)}

}

func (d Deck) GetSize() int {
	return len(d.cards)
}

func (d *Deck) Draw() (Card, error) {
	if len(d.cards) == 0 {
		return nil, errors.New("No cards left in deck")
	}
	topDeck := d.cards[0]
	d.cards = d.cards[1:]
	return topDeck, nil
}
