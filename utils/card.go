package utils

type CardSuit map[int]string
type CardNumber map[int]string

type Card struct {
	Suit   int
	Number int
}

func CreateCard(inputNumber int, inputColor int) Card {
	return Card{Number: inputNumber, Suit: inputColor}
}

func (card Card) HasSameNumber(anotherCard Card) bool {
	return card.Number == anotherCard.Number
}

func (card Card) HasSameSuit(anotherCard Card) bool {
	return card.Suit == anotherCard.Suit
}

func GetSuitMap() CardSuit {
	return CardSuit{
		CLUBS:    "c",
		DIAMONDS: "d",
		HEARTS:   "h",
		SPADES:   "s",
	}
}

func GetNumberMap() CardSuit {
	return CardSuit{
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "10",
		11: "J",
		12: "Q",
		13: "K",
		14: "A",
	}
}
