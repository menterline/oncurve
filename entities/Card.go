package entities

/*
Every thing is a card, but the main difference is a land or spell
*/
type Card interface {
	GetManaCost() int
	GetCardType() CardType
	CanPlay(boardState BoardState) bool
	Play(boardState *BoardState)
}
