package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerNew(t *testing.T) {
	player := NewPlayer()

	assert.Equal(t, player.Health, 20)
	assert.Len(t, player.Deck.Stack, 6)
	assert.Len(t, player.Deck.Hand, 0)
	assert.Len(t, player.Deck.Discard, 0)
}
