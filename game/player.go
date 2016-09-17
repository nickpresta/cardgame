package game

import "fmt"

// Player represents an entity playing a CardGame
type Player interface {
	UseCards(number int) []Card
	ReceiveCards(cards []Card)
	Hand() []Card
}

// DefaultPlayer is an implementation of a Player
type DefaultPlayer struct {
	Name string
	hand []Card
}

func (p *DefaultPlayer) String() string {
	return fmt.Sprintf("<Player Name=%s>", p.Name)
}

// ReceiveCards accepts a new set of cards to store in the Player's hand
func (p *DefaultPlayer) ReceiveCards(cards []Card) {
	p.hand = append(p.hand, cards...)
}

// UseCards plays the first number cards from a Player's hand
func (p *DefaultPlayer) UseCards(number int) []Card {
	// Use the first number cards from a hand
	cards := make([]Card, number)
	// shift from top, shrink slice
	cards, p.hand = p.hand[0:number], p.hand[number:]
	return cards
}

// Hand returns the Player's hand at this moment in time
func (p *DefaultPlayer) Hand() []Card {
	return p.hand
}
