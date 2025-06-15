package entities

import "errors"

// Not really "mana" pool, but resources to create mana
type ManaPool struct {
	AvailableMana map[Mana]int
}

func (m *ManaPool) GetMana() map[Mana]int {
	return m.AvailableMana
}

func (m *ManaPool) AddtoPool(mana Mana, amount int) {
	if m.AvailableMana == nil {
		m.AvailableMana = make(map[Mana]int)
	}
	m.AvailableMana[mana] += amount
}

func (m *ManaPool) RemoveFromPool(mana Mana) error {
	if m.AvailableMana == nil {
		return errors.New("mana pool is empty")
	}
	m.AvailableMana[mana] = m.AvailableMana[mana] - 1
	if m.AvailableMana[mana] <= 0 {
		delete(m.AvailableMana, mana)
	}
	return nil
}

func (m *ManaPool) CastSpell(spell Spell) {
	for _, mana := range spell.GetManaCost() {
		m.RemoveFromPool(mana)
	}
}
