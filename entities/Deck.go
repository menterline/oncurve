package entities

import (
	"errors"
	"math/rand"
	"time"
)

type Deck struct {
	cards []Card
}

const MAX_DECK_SIZE = 60

/*
when building the deck, we need a list of spells, and a list of lands,
then randomly merge them together
*/
func NewDeck(s SimSettingsData) Deck {
	spells := s.GetListOfSpells()
	lands := make([]BasicLand, MAX_DECK_SIZE-len(spells))
	tempDeck := make([]Card, 0, MAX_DECK_SIZE)
	for _, spell := range spells {
		tempDeck = append(tempDeck, spell)
	}
	for _, land := range lands {
		tempDeck = append(tempDeck, land)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(tempDeck), func(i, j int) { tempDeck[i], tempDeck[j] = tempDeck[j], tempDeck[i] })
	return Deck{cards: tempDeck}

}

func (d Deck) GetSize() int {
	return len(d.cards)
}

func (d Deck) GetCards() []Card {
	return d.cards
}

func (d *Deck) Draw() (Card, error) {
	if len(d.cards) == 0 {
		return nil, errors.New("No cards left in deck")
	}
	topDeck := d.cards[0]
	d.cards = d.cards[1:]
	return topDeck, nil
}
