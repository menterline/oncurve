package entities

type Spell struct {
	Cost     []Mana
	CardType CardType
}

func (s *Spell) GetManaCost() []Mana {
	return s.Cost
}

func (s *Spell) GetCardType() CardType {
	return s.CardType
}

func (s Spell) CanCast(manaPool ManaPool) bool {
	return manaPool.HasMana(s.Cost)
}
