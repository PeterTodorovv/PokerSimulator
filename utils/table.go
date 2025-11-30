package utils

import (
	"fmt"
	"strings"
)

type Table struct {
	deck         Deck
	players      []Hand
	river        []Card
	riverIsDrawn bool
}

func CreateTable(players int) Table {
	table := Table{deck: CreateDeck()}

	for i := 0; i < players; i++ {
		hand := Hand{}
		for j := 0; j < HAND_SIZE; j++ {
			hand.AddCard(table.deck.DrawCard())
		}
		table.players = append(table.players, hand)
	}

	return table
}

func (table Table) GetPlayers() []Hand {
	return table.players
}

func (table Table) GetRiver() []Card {
	return table.river
}

func (table Table) ExportRiver(length int) string {

	numberMap := GetNumberMap()
	suitMap := GetSuitMap()

	riverString := ""
	for i := 0; i < length; i++ {
		riverString += fmt.Sprintf("%s%s ", numberMap[table.river[i].Number], suitMap[table.river[i].Suit])
	}
	return strings.TrimSuffix(riverString, " \t\n")
}
func (table *Table) DrawCardToRiver() {
	if len(table.river) == 5 {
		panic("Trying to draw over 5 cards on the river")
	}

	if !table.riverIsDrawn {
		panic("Trying to draw a single card to the river before drawing the river")
	}

	table.river = append(table.river, table.deck.DrawCard())
}

func (table *Table) DrawRiver() {
	if table.riverIsDrawn {
		panic("River is already drawn!")
	}
	table.river = append(table.river, table.deck.DrawCard())
	table.river = append(table.river, table.deck.DrawCard())
	table.river = append(table.river, table.deck.DrawCard())
	table.riverIsDrawn = true
}
