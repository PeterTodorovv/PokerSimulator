package utils

import "fmt"

type Hand struct {
	cards []Card
}

func (hand *Hand) AddCard(card Card) {
	hand.cards = append(hand.cards, card)
}

func (hand Hand) GetCards() []Card {
	return hand.cards
}

func (hand Hand) ExportHand() string {
	cards := hand.GetCards()

	if cards[0].Number < cards[1].Number {
		tmp := cards[0]
		cards[0] = cards[1]
		cards[1] = tmp
	}

	if cards[0].Number == cards[1].Number && cards[0].Suit < cards[1].Suit {
		tmp := cards[0]
		cards[0] = cards[1]
		cards[1] = tmp
	}

	numberMap := GetNumberMap()
	suitMap := GetSuitMap()

	return fmt.Sprintf("%s%s %s%s", numberMap[cards[0].Number], suitMap[cards[0].Suit], numberMap[cards[1].Number], suitMap[cards[1].Suit])
}
