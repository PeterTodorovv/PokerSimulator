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
			fmt.Printf("Games: %d\n", i)
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

func SimulatePreFlopThreePlayers(simulationSize int) {
	db, err := sql.Open("sqlite", "simulations.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ThreePlayersPreFlop (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		hand TEXT,
		hits INTEGER,
		wins INTEGER,
		draws INTEGER,
		winChance REAL,
		drawChance REAL
	)`)

	handStats := map[string][]int{}
	playersCount := 3
	for i := 1; i <= simulationSize; i++ {
		game := utils.CreateGame(playersCount)
		for !game.IsGameOver() {
			game.AdvanceGame()
		}

		if i%10000000 == 0 {
			fmt.Printf("Games: %d\n", i)
		}
		hands := game.GetHands()

		for _, hand := range hands {
			_, ok := handStats[hand.ExportHand()]

			if !ok {
				handStats[hand.ExportHand()] = []int{0, 0, 0}
			}
			handStats[hand.ExportHand()][0]++
		}

		result := utils.GameResult{Winner: -1}
		currentWinner := 0
		for i := 0; i < playersCount-1; i++ {
			result = utils.CompareHands(hands[currentWinner], hands[i+1], game.GetRiver())
			if result.Winner == 0 {
				result.Winner = i
				currentWinner = i
			} else if result.Winner == 1 {
				result.Winner = i + 1
				currentWinner = i + 1
			}
		}

		if result.Winner == -1 {
			for _, hand := range hands {
				handStats[hand.ExportHand()][2]++
			}
		} else {
			handStats[hands[result.Winner].ExportHand()][1]++
		}
	}

	for key, stats := range handStats {
		_, err = db.Exec(`INSERT INTO ThreePlayersPreFlop (hand, hits, wins, draws, winChance, drawChance) VALUES (?, ?, ?, ?, ?, ?)`, key, stats[0], stats[1], stats[2], float32(stats[1])/float32(stats[0])*100, float32(stats[2])/float32(stats[0])*100)
	}
}

func SimulateCombinationChance(simulationSize int) {
	db, err := sql.Open("sqlite", "simulations.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Combinations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		hand TEXT,
		hits INTEGER,
		chance REAL
	)`)

	handStats := map[int]int{}
	playersCount := 1
	for i := 1; i <= simulationSize; i++ {
		game := utils.CreateGame(playersCount)
		for !game.IsGameOver() {
			game.AdvanceGame()
		}

		if i%10000000 == 0 {
			fmt.Printf("Games: %d\n", i)
		}
		hands := game.GetHands()

		result := utils.CalculateHand(hands[0], game.GetRiver())
		handStats[result.Combination]++
	}

	combinationMap := map[int]string{
		0 : "High Card",
		1 : "Pair",
		2 : "Two Pair",
		3 : "Three Of A Kind",
		4 : "Straight",
		5 : "Flush",
		6 : "Full House",
		7 : "Four Of A Kind",
		8 : "Straight Flush",
		9 : "High Card",

	}

	for key := 0; key < len(handStats);key++{
		hits := handStats[key]
		_, err = db.Exec(`INSERT INTO Combinations (hand, hits, chance) VALUES (?, ?, ?)`, combinationMap[key], hits, float32(hits) / float32(simulationSize) * 100)

	}
}
