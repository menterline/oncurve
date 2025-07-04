package entities

// Ensure that the Mana type is defined in a separate package (e.g., "types")
// and imported here, or move shared types to a new package to avoid import cycles.
// import "yourmodule/types" // Uncomment and adjust as needed

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
