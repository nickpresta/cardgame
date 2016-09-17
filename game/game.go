package game

type CardGame interface {
	Start()
	TakeTurn() TurnResult
	IsOver() bool
	Result() GameResult
}

// Result represents the outcome of a Turn in a CardGame
type TurnResult struct {
	Winners []Player
}

// Result represents the outcome of a Turn in a CardGame
type GameResult struct {
	Winners []Player
}
