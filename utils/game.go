package utils

type Game struct {
	table Table
	stage int
}

func CreateGame(playersCount int) Game {
	game := Game{table: CreateTable(playersCount), stage: 0}
	return game
}

func (game *Game) AdvanceGame() {
	if game.stage == 0 {
		game.table.DrawRiver()
	} else {
		game.table.DrawCardToRiver()
	}

	game.stage++
}

func (game Game) IsGameOver() bool {
	return game.stage == MAX_GAME_STAGE
}

func (game Game) GetHands() []Hand {
	return game.table.GetPlayers()
}

func (game Game) GetRiver() []Card {
	return game.table.GetRiver()
}

func (game Game) GetRiverString(length int) string {
	return game.table.ExportRiver(length)
}
