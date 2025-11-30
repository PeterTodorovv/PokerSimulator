package utils

import (
	"math/rand"
)

type Deck struct {
	cards []Card
}

func CreateDeck() Deck {
	deck := Deck{}
	for i := 0; i < COLORS_COUNT; i++ {
		for j := 2; j <= CARDS_PER_COLOR+1; j++ {
			deck.addCard(CreateCard(j, i))
		}
	}

	deck.shuffle()
	deck.shuffle()
	deck.shuffle()

	return deck
}

func (deck *Deck) DrawCard() Card {
	if len(deck.cards) == 0 {
		panic("Trying to draw card from an empty deck!")
	}

	drawnCard := deck.cards[0]
	deck.cards = deck.cards[1:]

	return drawnCard
}

func (deck *Deck) addCard(card Card) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck) shuffle() {
	max := len(deck.cards)
	counter := 0

	for counter < max {
		newPosition := rand.Intn(max)
		tmp := deck.cards[newPosition]
		deck.cards[newPosition] = deck.cards[counter]
		deck.cards[counter] = tmp
		counter++
	}
}
