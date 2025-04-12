package analyzer

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
)

var (
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
)

var characterNames = map[int]string{
	1: "Fox", 2: "Captain Falcon", 3: "Donkey Kong", 4: "Kirby", 5: "Bowser",
	6: "Link", 7: "Sheik", 8: "Ness", 9: "Peach", 10: "Ice Climbers", 11: "Nana", 12: "Pikachu",
	13: "Samus", 14: "Yoshi", 15: "Jigglypuff", 16: "Mewtwo", 17: "Luigi", 18: "Marth",
	19: "Zelda", 20: "Young Link", 21: "Dr. Mario", 22: "Falco", 23: "Pichu",
	24: "Mr. Game & Watch", 25: "Ganondorf", 26: "Roy", 0: "Mario",
	27: "Wolf", 28: "Diddy Kong", 29: "Charizard", 30: "Lucas", 31: "Sonic",
}

func PrintCharacterStats(stats map[int]*CharacterPlaytime) {
	type charStat struct {
		ID    int
		Stats *CharacterPlaytime
	}
	var list []charStat
	for id, s := range stats {
		if id >= 27 && id <= 31 {
			continue
		}
		list = append(list, charStat{id, s})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Stats.GameCount > list[j].Stats.GameCount
	})

	fmt.Println(blue("\n=== Playtime for Each Character ==="))
	for i, stat := range list {
		id := stat.ID
		p := stat.Stats
		name := characterNames[id]
		if name == "" {
			name = fmt.Sprintf("Unknown (ID: %d)", id)
		}
		totalSeconds := float64(p.FramesPlayed) / 60
		totalMinutes := totalSeconds / 60

		fmt.Println("----------------------------------")
		fmt.Println(rainbowText(i, len(list), "Character:", name))
		if totalMinutes >= 60 {
			fmt.Printf("%s %.2f hours\n", green("Time:"), totalMinutes/60)
		} else {
			fmt.Printf("%s %d minutes, %d seconds\n", green("Time:"), int(totalMinutes), int(totalSeconds)%60)
		}
		fmt.Printf("%s %d\n", yellow("Games:"), p.GameCount)
	}
}

func PrintTotalStats(games int, frames int) {
	totalSeconds := float64(frames) / 60
	totalMinutes := totalSeconds / 60
	totalHours := totalSeconds / 3600
	totalDays := totalSeconds / 86400

	fmt.Println(blue("\n=== Total Summary ==="))
	fmt.Printf("%s %d\n", magenta("Total Games Played:"), games)
	fmt.Print(magenta("Total Time Played: "))

	if totalDays >= 1 {
		fmt.Printf("%.2f hours (%.2f days)\n", totalHours, totalDays)
	} else if totalHours >= 1 {
		fmt.Printf("%.2f hours\n", totalHours)
	} else {
		fmt.Printf("%d minutes, %d seconds\n", int(totalMinutes), int(totalSeconds)%60)
	}
}
