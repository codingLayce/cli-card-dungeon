package entities

import (
	"cli-dungeon/game/deck"
)

type Player struct {
	Deck deck.Deck
	Entity
}

func NewPlayer() Player {
	return Player{
		Deck: deck.NewDeck([]deck.Card{
			deck.NewCard(1, "Frappe", 3, 0),
			deck.NewCard(2, "Frappe", 3, 0),
			deck.NewCard(3, "Frappe", 3, 0),
			deck.NewCard(4, "Défense", 0, 3),
			deck.NewCard(5, "Défense", 0, 3),
			deck.NewCard(6, "Défense", 0, 3),
		}),
		Entity: Entity{
			Health: 20,
		},
	}
}
