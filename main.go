package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nickpresta/cardgame/game"
	"github.com/nickpresta/cardgame/rulesets"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	deck := game.NewDeck()
	p1 := &game.DefaultPlayer{
		Name: "Nick",
	}
	p2 := &game.DefaultPlayer{
		Name: "Debo",
	}
	warGame := rulesets.NewWar(deck, []game.Player{p1, p2})
	warGame.Start()

	for {
		warGame.TakeTurn()

		if warGame.IsOver() {
			break
		}
	}
	fmt.Printf("Winner is: %s\n", warGame.Result().Winners[0])
}
