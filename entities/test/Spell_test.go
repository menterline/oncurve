package entities

import (
	"testing"

	"github.com/menterline/oncurve-service/entities"
)

func TestGetSpellReturnsAllMana(t *testing.T) {
	manaCost := []entities.Mana{entities.Black, entities.Black, entities.Blue}
	mySpell := entities.Spell{Cost: manaCost, CardType: entities.Creature}
	got := mySpell.GetManaCost()
	if len(got) != 3 {
		t.Errorf("expected mana cost length 3, got %d", len(got))
	}
	if mySpell.GetCardType() != entities.Creature {
		t.Errorf("expected card type Creature, got %v", mySpell.GetCardType())
	}
}
