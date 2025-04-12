package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/fatih/color"
	goslippi "github.com/pmcca/go-slippi"
	"github.com/sqweek/dialog"
)

var (
	cyan   = color.New(color.FgCyan).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
)

type CharacterPlaytime struct {
	CharacterID  int
	FramesPlayed int
	GameCount    int
	ReplayFiles  []string
}

func processSlippiFile(path string, wg *sync.WaitGroup, framesChan chan CharacterPlaytime, slippiCode string) {
	defer wg.Done()

	meta, err := goslippi.ParseMeta(path)
	if err != nil {
		return
	}

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return
	}

	var parsedMeta map[string]interface{}
	if err := json.Unmarshal(metaJSON, &parsedMeta); err != nil {
		return
	}

	players, ok := parsedMeta["Players"].(map[string]interface{})
	if !ok {
		return
	}

	for _, playerData := range players {
		player, ok := playerData.(map[string]interface{})
		if !ok {
			continue
		}

		names, _ := player["Names"].(map[string]interface{})
		if names["SlippiCode"] == slippiCode || names["Name"] == slippiCode {
			characters, _ := player["Characters"].([]interface{})
			if len(characters) > 0 {
				firstCharacter := characters[0].(map[string]interface{})
				characterID := int(firstCharacter["CharacterID"].(float64))
				framesPlayed := int(firstCharacter["FramesPlayed"].(float64))

				framesChan <- CharacterPlaytime{CharacterID: characterID, FramesPlayed: framesPlayed, GameCount: 1, ReplayFiles: []string{path}}
			}
			break
		}
	}
}

func main() {
	var slippiCode string
	fmt.Print("Enter your Slippi code or name: ")
	fmt.Scanln(&slippiCode)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %v", err)
	}
	defaultDir := filepath.Join(homeDir, "Documents", "Slippi")
	dir, err := dialog.Directory().Title("Select Your Slippi Replay Folder").SetStartDir(defaultDir).Browse()
	if err != nil {
		log.Fatalf("Error selecting folder: %v", err)
	}

	var wg sync.WaitGroup
	framesChan := make(chan CharacterPlaytime)
	totalGames := 0
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".slp" {
			totalGames++
			wg.Add(1)
			go processSlippiFile(path, &wg, framesChan, slippiCode)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking through the directory: %v", err)
	}

	go func() {
		wg.Wait()
		close(framesChan)
	}()

	characterPlaytime := make(map[int]*CharacterPlaytime)
	totalFrames := 0

	for playtime := range framesChan {
		if _, exists := characterPlaytime[playtime.CharacterID]; !exists {
			characterPlaytime[playtime.CharacterID] = &CharacterPlaytime{CharacterID: playtime.CharacterID}
		}
		characterPlaytime[playtime.CharacterID].FramesPlayed += playtime.FramesPlayed
		characterPlaytime[playtime.CharacterID].GameCount += playtime.GameCount
		characterPlaytime[playtime.CharacterID].ReplayFiles = append(characterPlaytime[playtime.CharacterID].ReplayFiles, playtime.ReplayFiles...)
		totalFrames += playtime.FramesPlayed
	}

	characterNames := map[int]string{
		1: "Fox", 2: "Captain Falcon", 3: "Donkey Kong", 4: "Kirby", 5: "Bowser",
		6: "Link", 7: "Sheik", 8: "Ness", 9: "Peach", 10: "Ice Climbers", 11: "Nana", 12: "Pikachu",
		13: "Samus", 14: "Yoshi", 15: "Jigglypuff", 16: "Mewtwo", 17: "Luigi", 18: "Marth",
		19: "Zelda", 20: "Young Link", 21: "Dr. Mario", 22: "Falco", 23: "Pichu",
		24: "Mr. Game & Watch", 25: "Ganondorf", 26: "Roy", 0: "Mario",
		27: "Wolf", 28: "Diddy Kong", 29: "Charizard", 30: "Lucas", 31: "Sonic",
	}

	var sortedIDs []int
	for characterID := range characterPlaytime {
		sortedIDs = append(sortedIDs, characterID)
	}
	sort.Ints(sortedIDs)

	const (
		Reset  = "\033[0m"
		Cyan   = "\033[36m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
	)

	fmt.Println("\n=== Playtime for Each Character ===")

	for _, characterID := range sortedIDs {
		playtime := characterPlaytime[characterID]
		charName := characterNames[characterID]
		if charName == "" {
			charName = fmt.Sprintf("Unknown (ID: %d)", characterID)
		}

		totalSeconds := float64(playtime.FramesPlayed) / 60
		totalMinutes := totalSeconds / 60

		fmt.Println("----------------------------------")
		fmt.Printf("%s %s\n", cyan("Character:"), charName)
		if totalMinutes >= 60 {
			fmt.Printf("%s %.2f hours\n", green("Time:"), totalMinutes/60)
		} else {
			fmt.Printf("%s %d minutes, %d seconds\n", green("Time:"), int(totalMinutes), int(totalSeconds)%60)
		}
		fmt.Printf("%s %d\n", yellow("Games:"), playtime.GameCount)
	}
	fmt.Scanln()
}
