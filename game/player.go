package game

type Player interface {
	ReceiveCards(cards []*Card)
	Hand() []*Card
}

// Player represents an entity playing a CardGame
type DefaultPlayer struct {
	Name string
	hand []*Card
}

func (p *DefaultPlayer) ReceiveCards(cards []*Card) {
	p.hand = append(p.hand, cards...)
}

func (p *DefaultPlayer) Hand() []*Card {
	return p.hand
}
