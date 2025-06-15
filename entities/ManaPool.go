package entities

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
