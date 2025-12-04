package tests

import (
	"PokerSimulator/utils"
	"testing"
)

func TestHighCardVsPair(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 1, Number: 14})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 4})
	hand2.AddCard(utils.Card{Suit: 1, Number: 10})

	river := []utils.Card{
		{Suit: 2, Number: 3},
		{Suit: 3, Number: 6},
		{Suit: 0, Number: 9},
		{Suit: 2, Number: 5},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2}, river)

	if result.Winner != 1 {
		t.Errorf("Expected winner 1; Got %d", result.Winner)
	}
}

func TestHighCardVsHighCard(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 1, Number: 14})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 4})
	hand2.AddCard(utils.Card{Suit: 1, Number: 13})

	river := []utils.Card{
		{Suit: 2, Number: 3},
		{Suit: 3, Number: 6},
		{Suit: 0, Number: 9},
		{Suit: 2, Number: 5},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2}, river)

	if result.Winner != 0 {
		t.Errorf("Expected winner 0; Got %d", result.Winner)
	}
}

func TestHighCardDraw(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 4})
	hand.AddCard(utils.Card{Suit: 1, Number: 14})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 4})
	hand2.AddCard(utils.Card{Suit: 2, Number: 14})

	river := []utils.Card{
		{Suit: 2, Number: 3},
		{Suit: 3, Number: 8},
		{Suit: 0, Number: 11},
		{Suit: 2, Number: 12},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2}, river)

	if result.Winner != -1 {
		t.Errorf("Expected winner -1; Got %d", result.Winner)
	}
}

func TestHighCardDraw2(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 1, Number: 14})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 4})
	hand2.AddCard(utils.Card{Suit: 2, Number: 14})

	river := []utils.Card{
		{Suit: 2, Number: 3},
		{Suit: 3, Number: 8},
		{Suit: 0, Number: 11},
		{Suit: 2, Number: 12},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2}, river)

	if result.Winner != -1 {
		t.Errorf("Expected winner -1; Got %d", result.Winner)
	}
}

func TestHighCardDraw3(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 1, Number: 3})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 4})
	hand2.AddCard(utils.Card{Suit: 2, Number: 5})

	river := []utils.Card{
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 8},
		{Suit: 0, Number: 11},
		{Suit: 2, Number: 12},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2}, river)

	if result.Winner != -1 {
		t.Errorf("Expected winner -1; Got %d", result.Winner)
	}
}

func TestHighCardDrawThreeHands(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 1, Number: 3})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 4})
	hand2.AddCard(utils.Card{Suit: 2, Number: 5})

	hand3 := utils.Hand{}
	hand3.AddCard(utils.Card{Suit: 1, Number: 4})
	hand3.AddCard(utils.Card{Suit: 3, Number: 5})

	river := []utils.Card{
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 8},
		{Suit: 0, Number: 11},
		{Suit: 2, Number: 12},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2, hand3}, river)

	if result.Winner != -1 {
		t.Errorf("Expected winner -1; Got %d", result.Winner)
	}
}

func TestThreeHandsWinFirst(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 12})
	hand.AddCard(utils.Card{Suit: 1, Number: 3})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 11})
	hand2.AddCard(utils.Card{Suit: 2, Number: 5})

	hand3 := utils.Hand{}
	hand3.AddCard(utils.Card{Suit: 1, Number: 4})
	hand3.AddCard(utils.Card{Suit: 3, Number: 5})

	river := []utils.Card{
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 8},
		{Suit: 0, Number: 11},
		{Suit: 2, Number: 12},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2, hand3}, river)

	if result.Winner != 0 {
		t.Errorf("Expected winner 0; Got %d", result.Winner)
	}
}

func TestThreeHandsWinSecond(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 12})
	hand.AddCard(utils.Card{Suit: 1, Number: 3})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 14})
	hand2.AddCard(utils.Card{Suit: 2, Number: 5})

	hand3 := utils.Hand{}
	hand3.AddCard(utils.Card{Suit: 1, Number: 4})
	hand3.AddCard(utils.Card{Suit: 3, Number: 5})

	river := []utils.Card{
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 8},
		{Suit: 0, Number: 11},
		{Suit: 2, Number: 12},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2, hand3}, river)

	if result.Winner != 1 {
		t.Errorf("Expected winner 1; Got %d", result.Winner)
	}
}

func TestThreeHandsWinThird(t *testing.T) {
	hand := utils.Hand{}
	hand.AddCard(utils.Card{Suit: 0, Number: 12})
	hand.AddCard(utils.Card{Suit: 1, Number: 3})

	hand2 := utils.Hand{}
	hand2.AddCard(utils.Card{Suit: 0, Number: 4})
	hand2.AddCard(utils.Card{Suit: 2, Number: 5})

	hand3 := utils.Hand{}
	hand3.AddCard(utils.Card{Suit: 1, Number: 14})
	hand3.AddCard(utils.Card{Suit: 3, Number: 5})

	river := []utils.Card{
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 8},
		{Suit: 0, Number: 11},
		{Suit: 2, Number: 12},
		{Suit: 3, Number: 10},
	}

	result := utils.CompareHands([]utils.Hand{hand, hand2, hand3}, river)

	if result.Winner != 2 {
		t.Errorf("Expected winner 2; Got %d", result.Winner)
	}
}