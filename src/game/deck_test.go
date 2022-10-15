package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultCards = CardStack{
	NewCard(1, "Frappe", 3, 0),
	NewCard(2, "Frappe", 3, 0),
	NewCard(3, "Frappe", 3, 0),
	NewCard(4, "Défense", 0, 3),
	NewCard(5, "Défense", 0, 3),
	NewCard(6, "Défense", 0, 3),
}

func TestNewDeck(t *testing.T) {
	deck := NewDeck(defaultCards)

	assert.Len(t, deck.Stack, 6)
	assert.Len(t, deck.Hand, 0)
	assert.Len(t, deck.Discard, 0)
}

func TestDeck_ShuffleStack(t *testing.T) {
	deck := NewDeck(defaultCards)
	deck.ShuffleStack()

	assert.Len(t, deck.Stack, 6)
	assert.Len(t, deck.Hand, 0)
	assert.Len(t, deck.Discard, 0)

	// Check that no duplication has appeared
	cache := make(map[int]struct{})
	for _, card := range deck.Stack {
		_, present := cache[card.ID]
		assert.False(t, present, "The card with ID (%d) is present multiple times", card.ID)
		cache[card.ID] = struct{}{}
	}

	// Check that at least one card's index has changed
	atLeastOneChanged := false
	for i, card := range deck.Stack {
		if card.ID != i+1 {
			atLeastOneChanged = true
		}
	}
	assert.True(t, atLeastOneChanged, "All cards are at the same index than before shuffling")

}

func TestDeck_InitialDraw(t *testing.T) {
	deck := NewDeck(defaultCards)
	deck.ShuffleStack()
	deck.InitialDraw()

	assert.Len(t, deck.Stack, 4)
	assert.Len(t, deck.Hand, 2)
	assert.Len(t, deck.Discard, 0)
}

func TestDeck_PlayCard(t *testing.T) {
	deck := NewDeck(defaultCards)
	deck.ShuffleStack()
	deck.InitialDraw()

	err := deck.PlayCard(deck.Hand[0].ID)
	assert.NoError(t, err)

	assert.Len(t, deck.Hand, 1)
	assert.Len(t, deck.Stack, 4)
	assert.Len(t, deck.Discard, 1)
}

func TestDeck_PlayCard_Error(t *testing.T) {
	deck := NewDeck(defaultCards)
	deck.ShuffleStack()
	deck.InitialDraw()

	err := deck.PlayCard(89)
	assert.EqualError(t, err, ErrCardIDNotExists.Error())
}

func TestDeck_Draw(t *testing.T) {
	strikeCard1 := NewCard(1, "Frappe", 3, 0)
	strikeCard2 := NewCard(2, "Frappe", 3, 0)
	shieldCard1 := NewCard(3, "Défense", 0, 3)
	shieldCard2 := NewCard(4, "Défense", 0, 3)

	for name, tc := range map[string]struct {
		stackCards        CardStack
		discardStackCards CardStack
		nbToDraw          int
		expectedStackLen  int
		expectedHandLen   int
	}{
		"With enough cards in stack": {
			stackCards:       CardStack{strikeCard1, strikeCard2, shieldCard1, shieldCard2},
			nbToDraw:         2,
			expectedHandLen:  2,
			expectedStackLen: 2,
		},
		"With not enough cards in stack": {
			stackCards:        CardStack{strikeCard1},
			discardStackCards: CardStack{strikeCard2, shieldCard1, shieldCard2},
			nbToDraw:          2,
			expectedStackLen:  2,
			expectedHandLen:   2,
		},
	} {
		t.Run(name, func(t *testing.T) {
			deck := NewDeck(tc.stackCards)
			deck.ShuffleStack()
			deck.Discard = tc.discardStackCards

			deck.Draw(tc.nbToDraw)

			assert.Len(t, deck.Stack, tc.expectedStackLen)
			assert.Len(t, deck.Hand, tc.expectedHandLen)
		})
	}
}
