package entities

type CardType int

const (
	Creature CardType = iota
	Instant
	Sorcery
	Artifact
)
