package game

// CardGame represents a conceptual card game
type CardGame interface {
	Start()
	TakeTurn() TurnResult
	IsOver() bool
	Result() Result
}

// TurnResult represents the outcome of a Turn in a CardGame
type TurnResult struct {
	Winners []Player
}

// Result represents the outcome of a Turn in a CardGame
type Result struct {
	Winners []Player
}
