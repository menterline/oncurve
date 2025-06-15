package entities

type Land struct {
	ManaTypes map[Mana]int
}

func (s *Land) GetManaTypes() map[Mana]int {
	return s.ManaTypes
}

func LandFactory(manaTypes []Mana) Land {
	manaMap := make(map[Mana]int)
	for _, mana := range manaTypes {
		manaMap[mana] = manaMap[mana] + 1
	}
	return Land{ManaTypes: manaMap}
}
