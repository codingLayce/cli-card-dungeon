package game

import (
	"errors"
)

var ErrCardIDNotExists = errors.New("card ID doesn't exists in current stack")

type CardStack []Card

func (stack *CardStack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *CardStack) Push(card Card) {
	*stack = append(*stack, card)
}

func (stack *CardStack) PushBack(card Card) {
	if stack.IsEmpty() {
		*stack = append(*stack, card)
		return
	}

	tmp := *stack
	*stack = CardStack{card}
	*stack = append(*stack, tmp...)
}

func (stack *CardStack) Pop() Card {
	lastIndex := len(*stack) - 1
	topCard := (*stack)[lastIndex]
	*stack = (*stack)[:lastIndex]

	return topCard
}

func (stack *CardStack) PopByID(cardID int) (Card, error) {
	card := Card{}
	index := -1

	for i, c := range *stack {
		if cardID == c.ID {
			card = c
			index = i
			break
		}
	}

	if index == -1 {
		return Card{}, ErrCardIDNotExists
	}

	if index == len(*stack)-1 {
		*stack = (*stack)[:index]
		return card, nil
	}

	tmp := (*stack)[index+1:]
	*stack = (*stack)[:index]
	*stack = append(*stack, tmp...)

	return card, nil
}
