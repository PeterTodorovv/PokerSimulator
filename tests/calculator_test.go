package tests

import (
	"PokerSimulator/utils"
	"fmt"
	"slices"
	"testing"
)

func TestCalculatorHighCard(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 3},
		{Suit: 3, Number: 6},
		{Suit: 0, Number: 9},
		{Suit: 2, Number: 5},
		{Suit: 3, Number: 10},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 1, Number: 14})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 10, 9, 6, 5}

	if result.Combination != 0 {
		t.Errorf("Expected 0; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards")
	}
}

func TestCalculatorHighCardOnRiver(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 3},
		{Suit: 3, Number: 6},
		{Suit: 0, Number: 9},
		{Suit: 2, Number: 5},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 1, Number: 10})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 10, 9, 6, 5}

	if result.Combination != 0 {
		t.Errorf("Expected 0; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestCalculatorPair(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 6},
		{Suit: 0, Number: 9},
		{Suit: 2, Number: 5},
		{Suit: 3, Number: 10},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 1, Number: 14})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{2, 2, 14, 10, 9}

	if result.Combination != 1 {
		t.Errorf("Expected 1; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestCalculatorPairRiver(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 2},
		{Suit: 0, Number: 9},
		{Suit: 2, Number: 5},
		{Suit: 3, Number: 10},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 6})
	hand.AddCard(utils.Card{Suit: 1, Number: 14})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{2, 2, 14, 10, 9}

	if result.Combination != 1 {
		t.Errorf("Expected 1; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestCalculatorTwoPair(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 2},
		{Suit: 0, Number: 9},
		{Suit: 2, Number: 5},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 6})
	hand.AddCard(utils.Card{Suit: 1, Number: 14})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 14, 2, 2, 9}

	if result.Combination != 2 {
		t.Errorf("Expected 2; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestCalculatorTwoPairRiver(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 2},
		{Suit: 0, Number: 9},
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 6})
	hand.AddCard(utils.Card{Suit: 1, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 14, 2, 2, 9}

	if result.Combination != 2 {
		t.Errorf("Expected 2; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestCalculatorTwoPairWhenThreePairs(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 2},
		{Suit: 0, Number: 6},
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 6})
	hand.AddCard(utils.Card{Suit: 1, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 14, 6, 6, 5}

	if result.Combination != 2 {
		t.Errorf("Expected 2; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestCalculatorThreeOfAKind(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 6},
		{Suit: 0, Number: 6},
		{Suit: 2, Number: 12},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 6})
	hand.AddCard(utils.Card{Suit: 1, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{6, 6, 6, 14, 12}

	if result.Combination != 3 {
		t.Errorf("Expected 3; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestCalculatorThreeOfAKindRiver(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 6},
		{Suit: 0, Number: 6},
		{Suit: 2, Number: 6},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 12})
	hand.AddCard(utils.Card{Suit: 1, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{6, 6, 6, 14, 12}

	if result.Combination != 3 {
		t.Errorf("Expected 3; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestStraight(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 3},
		{Suit: 0, Number: 4},
		{Suit: 2, Number: 6},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 12})
	hand.AddCard(utils.Card{Suit: 1, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{6, 5, 4, 3, 2}

	if result.Combination != 4 {
		t.Errorf("Expected 4; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestStraightAceToFive(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 3},
		{Suit: 0, Number: 4},
		{Suit: 2, Number: 7},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 12})
	hand.AddCard(utils.Card{Suit: 1, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{5, 4, 3, 2, 14}

	if result.Combination != 4 {
		t.Errorf("Expected 4; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestStraightWithPair(t *testing.T) {
	hand := utils.Hand{}
	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 3},
		{Suit: 0, Number: 4},
		{Suit: 2, Number: 7},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 14})
	hand.AddCard(utils.Card{Suit: 1, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{5, 4, 3, 2, 14}

	if result.Combination != 4 {
		t.Errorf("Expected 4; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestFlush(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 0, Number: 8},
		{Suit: 0, Number: 4},
		{Suit: 2, Number: 7},
		{Suit: 0, Number: 10},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 14})
	hand.AddCard(utils.Card{Suit: 0, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 10, 8, 5, 4}

	if result.Combination != 5 {
		t.Errorf("Expected 5; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestFlushWithMoreThanFiveCards(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 0, Number: 2},
		{Suit: 0, Number: 8},
		{Suit: 0, Number: 4},
		{Suit: 0, Number: 7},
		{Suit: 0, Number: 10},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 14})
	hand.AddCard(utils.Card{Suit: 0, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 10, 8, 7, 5}

	if result.Combination != 5 {
		t.Errorf("Expected 5; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestFlushWithStraight(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 0, Number: 3},
		{Suit: 0, Number: 4},
		{Suit: 2, Number: 7},
		{Suit: 0, Number: 10},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 14})
	hand.AddCard(utils.Card{Suit: 0, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 10, 5, 4, 3}

	if result.Combination != 5 {
		t.Errorf("Expected 5; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestFullHouse(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 3},
		{Suit: 0, Number: 5},
		{Suit: 2, Number: 5},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 1, Number: 14})
	hand.AddCard(utils.Card{Suit: 0, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{5, 5, 5, 14, 14}

	if result.Combination != 6 {
		t.Errorf("Expected 6; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestFullHouseWithTwoSetsOfTrheeCards(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 14},
		{Suit: 0, Number: 5},
		{Suit: 2, Number: 5},
		{Suit: 0, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 1, Number: 14})
	hand.AddCard(utils.Card{Suit: 0, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 14, 14, 5, 5}

	if result.Combination != 6 {
		t.Errorf("Expected 6; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestFourOfAKind(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 14},
		{Suit: 0, Number: 6},
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 1, Number: 14})
	hand.AddCard(utils.Card{Suit: 0, Number: 5})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 14, 14, 14, 6}

	if result.Combination != 7 {
		t.Errorf("Expected 6; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestFourOfAKindWithFullHouse(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 2},
		{Suit: 3, Number: 3},
		{Suit: 0, Number: 6},
		{Suit: 2, Number: 3},
		{Suit: 3, Number: 3},
	}

	hand.AddCard(utils.Card{Suit: 1, Number: 3})
	hand.AddCard(utils.Card{Suit: 0, Number: 6})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{3, 3, 3, 3, 6}

	if result.Combination != 7 {
		t.Errorf("Expected 6; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestFourOfAKindWithFullHouseAndKicker(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 13},
		{Suit: 3, Number: 14},
		{Suit: 0, Number: 6},
		{Suit: 2, Number: 14},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 1, Number: 14})
	hand.AddCard(utils.Card{Suit: 0, Number: 6})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{14, 14, 14, 14, 13}

	if result.Combination != 7 {
		t.Errorf("Expected 6; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestStraightFlush(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 13},
		{Suit: 0, Number: 4},
		{Suit: 0, Number: 6},
		{Suit: 0, Number: 5},
		{Suit: 3, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 0, Number: 3})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{6, 5, 4, 3, 2}

	if result.Combination != 8 {
		t.Errorf("Expected 6; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}

func TestStraightFlushAceToFive(t *testing.T) {
	hand := utils.Hand{}

	river := []utils.Card{
		{Suit: 2, Number: 13},
		{Suit: 0, Number: 4},
		{Suit: 0, Number: 7},
		{Suit: 0, Number: 5},
		{Suit: 0, Number: 14},
	}

	hand.AddCard(utils.Card{Suit: 0, Number: 2})
	hand.AddCard(utils.Card{Suit: 0, Number: 3})

	result := utils.CalculateHand(hand, river)
	expectedActiveCards := []int{5, 4, 3, 2, 14}

	if result.Combination != 8 {
		t.Errorf("Expected 6; Got %d", result.Combination)
	}

	if !slices.Equal(result.ActicveCards, expectedActiveCards) {
		t.Errorf("Wrong active Cards. Expected: %s; Actual: %s", fmt.Sprint(expectedActiveCards), fmt.Sprint(result.ActicveCards))
	}
}