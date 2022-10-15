package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardStack_IsEmpty(t *testing.T) {
	for name, tc := range map[string]struct {
		stack    CardStack
		expected bool
	}{
		"Should be empty": {
			stack:    CardStack{},
			expected: true,
		},
		"Shouldn't be empty": {
			stack:    defaultCards,
			expected: false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.stack.IsEmpty())
		})
	}
}

func TestCardStack_Push(t *testing.T) {
	card := NewCard(1, "Frappe", 3, 0)
	card2 := NewCard(2, "Frappe", 3, 0)

	for name, tc := range map[string]struct {
		stack         CardStack
		card          Card
		expectedStack CardStack
	}{
		"Empty stack": {
			card:          card,
			expectedStack: CardStack{card},
		},
		"With cards stack": {
			stack:         CardStack{card},
			card:          card2,
			expectedStack: CardStack{card, card2},
		},
	} {
		t.Run(name, func(t *testing.T) {
			tc.stack.Push(tc.card)

			assert.Len(t, tc.stack, len(tc.expectedStack))
			for i, value := range tc.expectedStack {
				assert.Equal(t, value, tc.stack[i])
			}
		})
	}
}

func TestCardStack_PushBack(t *testing.T) {
	strikeCard := NewCard(1, "Frappe", 3, 0)
	shieldCard := NewCard(2, "Défense", 0, 3)

	for name, tc := range map[string]struct {
		stack         CardStack
		card          Card
		expectedStack CardStack
	}{
		"Empty stack": {
			card:          strikeCard,
			expectedStack: CardStack{strikeCard},
		},
		"With cards stack": {
			stack:         CardStack{strikeCard, strikeCard, strikeCard},
			card:          shieldCard,
			expectedStack: CardStack{shieldCard, strikeCard, strikeCard, strikeCard},
		},
	} {
		t.Run(name, func(t *testing.T) {
			tc.stack.PushBack(tc.card)

			assert.Len(t, tc.stack, len(tc.expectedStack))
			for i, value := range tc.expectedStack {
				assert.Equal(t, value, tc.stack[i])
			}
		})
	}
}

func TestCardStack_Pop(t *testing.T) {
	strikeCard := NewCard(1, "Frappe", 3, 0)
	shieldCard := NewCard(2, "Défense", 0, 3)

	stack := CardStack{strikeCard, shieldCard}
	card := stack.Pop()

	assert.Equal(t, shieldCard, card)
	assert.Len(t, stack, 1)
}

func TestCardStack_PopByID(t *testing.T) {
	strikeCard := NewCard(1, "Frappe", 3, 0)
	shieldCard := NewCard(2, "Défense", 0, 3)
	strikeCard2 := NewCard(3, "Frappe", 3, 0)
	shieldCard2 := NewCard(4, "Défense", 0, 3)
	strikeCard3 := NewCard(5, "Frappe", 3, 0)
	shieldCard3 := NewCard(6, "Défense", 0, 3)

	for name, tc := range map[string]struct {
		cardID        int
		expectedStack CardStack
		expectedErr   string
	}{
		"Pop back card": {
			cardID:        1,
			expectedStack: CardStack{shieldCard, strikeCard2, shieldCard2, strikeCard3, shieldCard3},
		},
		"Pop top card": {
			cardID:        6,
			expectedStack: CardStack{strikeCard, shieldCard, strikeCard2, shieldCard2, strikeCard3},
		},
		"Pop other card": {
			cardID:        4,
			expectedStack: CardStack{strikeCard, shieldCard, strikeCard2, strikeCard3, shieldCard3},
		},
	} {
		t.Run(name, func(t *testing.T) {
			stack := CardStack{strikeCard, shieldCard, strikeCard2, shieldCard2, strikeCard3, shieldCard3}

			card, err := stack.PopByID(tc.cardID)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			assert.Equal(t, tc.cardID, card.ID)

			assert.Len(t, stack, len(tc.expectedStack))
			for i, value := range tc.expectedStack {
				assert.Equal(t, value, stack[i])
			}
		})
	}
}
