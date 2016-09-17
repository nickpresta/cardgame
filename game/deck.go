package game

import "fmt"

// CardSuit represents a Card suit in our Game
type CardSuit string

// Suit represents a specific set of suits in a specific Card Game
var Suit = struct {
	Heart   CardSuit
	Diamond CardSuit
	Spade   CardSuit
	Club    CardSuit
}{
	"Heart", "Diamond", "Spade", "Club",
}

func defaultSuits() [4]CardSuit {
	return [4]CardSuit{Suit.Diamond, Suit.Heart, Suit.Club, Suit.Spade}
}

// CardFace represents a card face value in our Game
type CardFace string

// Face represents a specific set of faces in a specific Card Game
var Face = struct {
	Two   CardFace
	Three CardFace
	Four  CardFace
	Five  CardFace
	Six   CardFace
	Seven CardFace
	Eight CardFace
	Nine  CardFace
	Ten   CardFace
	Jack  CardFace
	Queen CardFace
	King  CardFace
	Ace   CardFace
}{
	"Two", "Three", "Four", "Five",
	"Six", "Seven", "Eight", "Nine",
	"Ten", "Jack", "Queen", "King",
	"Ace",
}

func defaultFaces() [13]CardFace {
	return [13]CardFace{
		Face.Two, Face.Three, Face.Four, Face.Five,
		Face.Six, Face.Seven, Face.Eight, Face.Nine,
		Face.Ten, Face.Jack, Face.Queen, Face.King,
		Face.Ace,
	}
}

type Card struct {
	Suit CardSuit
	Face CardFace
}

func (c *Card) String() string {
	return fmt.Sprintf("<Card suit=%q value=%q>", c.Suit, c.Face)
}

// GameDeck represents a deck that can dispense cards and tell if it's empty
type Deck interface {
	TakeCards(number int) []Card
	IsEmpty() bool
}

// Deck represents a collection of Card types
type DefaultDeck struct {
	cards []Card
}

// TakeCards removes number of Cards from the top of the Deck
func (d *DefaultDeck) TakeCards(number int) []Card {
	cards := make([]Card, number)
	// shift from top, shrink slice
	cards, d.cards = d.cards[0:(number-1)], d.cards[number:]
	return cards
}

func (d *DefaultDeck) IsEmpty() bool {
	return len(d.cards) == 0
}

func NewDeck() *DefaultDeck {
	cards := make([]Card, 52)
	c := 0
	for _, suit := range defaultSuits() {
		for _, face := range defaultFaces() {
			cards[c] = Card{
				Suit: suit,
				Face: face,
			}
			c = c + 1
		}
	}
	return &DefaultDeck{
		cards: cards,
	}
}
