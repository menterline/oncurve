package entities

import (
	"errors"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

const MAX_DECK_SIZE = 60

/*
when building the deck, we need a list of spells, and a list of lands
combine them into a deck - NOT SHUFFLED
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
	return Deck{Cards: tempDeck}
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

func (d Deck) GetSize() int {
	return len(d.Cards)
}

func (d Deck) GetCards() []Card {
	return d.Cards
}

func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) == 0 {
		return nil, errors.New("No cards left in deck")
	}
	topDeck := d.Cards[0]
	d.Cards = d.Cards[1:]
	return topDeck, nil
}

func (d *Deck) DrawMultiple(num int) ([]Card, error) {
	cards := make([]Card, 0, num)
	for i := 0; i < num; i++ {
		card, err := d.Draw()
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
