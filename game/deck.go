package game

import (
	"fmt"
	"math/rand"
)

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

func defaultSuits() []CardSuit {
	return []CardSuit{Suit.Diamond, Suit.Heart, Suit.Club, Suit.Spade}
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

func faceAsInt(face CardFace) int {
	switch face {
	case Face.Two:
		return 2
	case Face.Three:
		return 3
	case Face.Four:
		return 4
	case Face.Five:
		return 5
	case Face.Six:
		return 6
	case Face.Seven:
		return 7
	case Face.Eight:
		return 8
	case Face.Nine:
		return 9
	case Face.Ten:
		return 10
	case Face.Jack:
		return 11
	case Face.Queen:
		return 12
	case Face.King:
		return 13
	case Face.Ace:
		return 14
	}
	return -1
}

func defaultFaces() []CardFace {
	return []CardFace{
		Face.Two, Face.Three, Face.Four, Face.Five,
		Face.Six, Face.Seven, Face.Eight, Face.Nine,
		Face.Ten, Face.Jack, Face.Queen, Face.King,
		Face.Ace,
	}
}

// Card is a representation of a Suit/Face pair
type Card struct {
	Suit CardSuit
	Face CardFace
}

func (c Card) String() string {
	return fmt.Sprintf("<Card suit=%q value=%q>", c.Suit, c.Face)
}

// Compare compares two cards against Face value
func (c Card) Compare(other Card) int {
	cInt := faceAsInt(c.Face)
	oInt := faceAsInt(other.Face)
	if cInt > oInt {
		return 1
	} else if cInt < oInt {
		return -1
	}
	return 0
}

// Deck represents a deck that can dispense cards and tell if it's empty
type Deck interface {
	TakeCards(number int) []Card
	IsEmpty() bool
	Shuffle()
}

// DefaultDeck represents a default implementation of a Deck
type DefaultDeck struct {
	cards []Card
}

// TakeCards removes number of Cards from the top of the Deck
func (d *DefaultDeck) TakeCards(number int) []Card {
	cards := make([]Card, number)
	// shift from top, shrink slice
	cards, d.cards = d.cards[0:number], d.cards[number:]
	return cards
}

// IsEmpty returns if the deck is empty
func (d *DefaultDeck) IsEmpty() bool {
	return len(d.cards) == 0
}

// Shuffle shuffles a deck of cards
func (d *DefaultDeck) Shuffle() {
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// NewDeck returns a new instance of a DefaultDeck
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
