package main

import (
	"fmt"

	"github.com/nickpresta/cardgame/game"
	"github.com/nickpresta/cardgame/rulesets"
)

func main() {
	deck := game.NewDeck()
	p1 := game.DefaultPlayer{
		Name: "Nick",
	}
	p2 := game.DefaultPlayer{
		Name: "Debo",
	}
	warGame := rulesets.War{
		Deck: deck,
		Players: []game.Player{
			p1, p2,
		},
	}

	warGame.Start()

	for {
		warGame.TakeTurn()

		if warGame.IsOver() {
			break
		}
	}
	fmt.Println(warGame.Result())
}
