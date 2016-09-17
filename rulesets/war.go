package rulesets

import (
	"fmt"

	"github.com/nickpresta/cardgame/game"
)

// War represents a specific Card Game
type War struct {
	handSize int
	deck     game.Deck
	players  []game.Player
	turnWins map[game.Player]int
}

// NewWar returns an instance of a War game
func NewWar(deck game.Deck, players []game.Player) *War {
	turnWins := make(map[game.Player]int)
	for _, player := range players {
		turnWins[player] = 0
	}

	return &War{
		handSize: 5,
		deck:     deck,
		players:  players,
		turnWins: turnWins,
	}
}

// Start starts off the game
func (w *War) Start() {
	w.deck.Shuffle()

	for _, p := range w.players {
		// Take w.handSize cards at the same time (don't alternative deal)
		p.ReceiveCards(w.deck.TakeCards(w.handSize))
	}
}

// IsOver determines if the CardGame is completely over
func (w *War) IsOver() bool {
	return w.deck.IsEmpty()
}

// TakeTurn takes a turn
func (w *War) TakeTurn() game.TurnResult {
	p1 := w.players[0]
	p2 := w.players[1]
	winner := w.battle(p1, p2)
	result := game.TurnResult{
		Winners: []game.Player{winner},
	}

	p1.ReceiveCards(w.deck.TakeCards(1))
	p2.ReceiveCards(w.deck.TakeCards(1))

	return result
}

func (w *War) battle(p1 game.Player, p2 game.Player) game.Player {
	p1Card := p1.UseCards(1)
	p2Card := p2.UseCards(1)

	compare := p1Card[0].Compare(p2Card[0])
	if compare > 0 {
		w.turnWins[p1] = w.turnWins[p1] + 1
		return p1
	} else if compare < 0 {
		w.turnWins[p2] = w.turnWins[p2] + 1
		return p2
	}
	// TODO: Take suit into account for no draw
	return p1
}

// Result gets the game's result at this moment in time
func (w *War) Result() game.Result {
	fmt.Printf("Turn Wins: %+v\n", w.turnWins)

	winner := w.players[0]
	for _, player := range w.players[1:] {
		if w.turnWins[player] > w.turnWins[winner] {
			winner = player
		}
	}

	return game.Result{
		Winners: []game.Player{winner},
	}
}
