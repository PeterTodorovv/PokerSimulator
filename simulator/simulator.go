package simulator

import (
	"PokerSimulator/utils"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func SimulatePreFlopStandOff(simulationSize int) {
	db, err := sql.Open("sqlite", "simulations.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS TwoPlayersPreFlop (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		hand TEXT,
		hits INTEGER,
		wins INTEGER,
		draws INTEGER,
		winChance REAL,
		drawChance REAL
	)`)

	handStats := map[string][]int{}
	playersCount := 2
	for i := 1; i <= simulationSize; i++ {
		game := utils.CreateGame(playersCount)
		for !game.IsGameOver() {
			game.AdvanceGame()
		}

		if i%10000000 == 0 {
			fmt.Printf("Spins: %d\n", i)
		}
		hands := game.GetHands()

		for _, hand := range hands {
			_, ok := handStats[hand.ExportHand()]

			if !ok {
				handStats[hand.ExportHand()] = []int{0, 0, 0}
			}
			handStats[hand.ExportHand()][0]++
		}

		result := utils.CompareHands(hands[0], hands[1], game.GetRiver())
		if result.Winner == -1 {
			for _, hand := range hands {
				handStats[hand.ExportHand()][2]++
			}
		} else {
			handStats[hands[result.Winner].ExportHand()][1]++
		}
	}

	for key, stats := range handStats {
		_, err = db.Exec(`INSERT INTO TwoPlayersPreFlop (hand, hits, wins, draws, winChance, drawChance) VALUES (?, ?, ?, ?, ?, ?)`, key, stats[0], stats[1], stats[2], float32(stats[1])/float32(stats[0])*100, float32(stats[2])/float32(stats[0])*100)
	}
}
