package main

import (
	"fmt"

	"github.com/9je/slippi-playtime-go/analyzer"
	"github.com/9je/slippi-playtime-go/utils"
)

func main() {
	slippiCode := utils.GetSlippiCode()
	dir := utils.PromptForDirectory()

	characterStats, totalGames, totalFrames := analyzer.Analyze(dir, slippiCode)
	analyzer.PrintCharacterStats(characterStats)
	analyzer.PrintTotalStats(totalGames, totalFrames)
	fmt.Scanln()
}
