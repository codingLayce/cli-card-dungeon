package deck

import (
	"math/rand"
	"time"
)

type Deck struct {
	Stack   CardStack
	Hand    CardStack
	Discard CardStack
}

func NewDeck(cards []Card) Deck {
	return Deck{
		Stack: cards,
	}
}

func (deck *Deck) ShuffleStack() {
	rand.Seed(time.Now().UnixMicro())
	rand.Shuffle(len(deck.Stack), func(i, j int) {
		deck.Stack[i], deck.Stack[j] = deck.Stack[j], deck.Stack[i]
	})
}

func (deck *Deck) InitialDraw() {
	deck.Hand.Push(deck.Stack.Pop())
	deck.Hand.Push(deck.Stack.Pop())
}

func (deck *Deck) PlayCard(cardID int) error {
	card, err := deck.Hand.PopByID(cardID)
	if err != nil {
		return err
	}

	deck.Discard.Push(card)

	return nil
}

func (deck *Deck) Draw(count int) {
	for i := 0; i < count; i++ {
		deck.Hand.Push(deck.Stack.Pop())
		if deck.Stack.IsEmpty() {
			deck.Stack = deck.Discard
			deck.Discard = CardStack{}
			deck.ShuffleStack()
		}
	}
}
