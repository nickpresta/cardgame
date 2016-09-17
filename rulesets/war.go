package rulesets

import "github.com/nickpresta/cardgame/game"

// War represents a specific Card Game
type War struct {
	handSize int
	Deck     *game.Deck
	Players  []game.Player
}

func (w *War) Start() {
	w.Deck.Shuffle()

	for _, p := range w.Players {
		// Take w.handSize cards at the same time (don't alternative deal)
		p.ReceiveCards(Deck.TakeCard(w.handSize))
	}
}

// IsOver determines if the CardGame is completely over
func (w *War) IsOver() bool {
	return w.Deck.IsEmpty()kk
}

func (w *War) TakeTurn() game.TurnResult {
	return game.TurnResult{
		Winners: []game.Player{},
	}
}

func (w *War) Result() game.GameResult {
	return game.GameResult{
		Winners: []game.Player{},
	}
}
