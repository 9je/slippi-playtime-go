package analyzer

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"

	goslippi "github.com/pmcca/go-slippi"
)

type CharacterPlaytime struct {
	CharacterID  int
	FramesPlayed int
	GameCount    int
	Wins         int
	Losses       int
	ReplayFiles  []string
}

func Analyze(dir, slippiCode string) (map[int]*CharacterPlaytime, int, int) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	characterPlaytime := make(map[int]*CharacterPlaytime)
	totalFrames := 0
	userGamesCount := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || filepath.Ext(path) != ".slp" {
			return nil
		}
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			if result := processSlippiFile(p, slippiCode); result != nil {
				mu.Lock()
				if _, exists := characterPlaytime[result.CharacterID]; !exists {
					characterPlaytime[result.CharacterID] = &CharacterPlaytime{CharacterID: result.CharacterID}
				}
				cp := characterPlaytime[result.CharacterID]
				cp.FramesPlayed += result.FramesPlayed
				cp.GameCount += result.GameCount
				cp.ReplayFiles = append(cp.ReplayFiles, result.ReplayFiles...)
				totalFrames += result.FramesPlayed
				userGamesCount++
				mu.Unlock()
			}
		}(path)
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking through the directory: %v", err)
	}

	wg.Wait()
	return characterPlaytime, userGamesCount, totalFrames
}

func processSlippiFile(path string, slippiCode string) *CharacterPlaytime {
	meta, err := goslippi.ParseMeta(path)
	if err != nil {
		return nil
	}

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return nil
	}

	var parsedMeta map[string]interface{}
	if err := json.Unmarshal(metaJSON, &parsedMeta); err != nil {
		return nil
	}

	players, ok := parsedMeta["Players"].(map[string]interface{})
	if !ok {
		return nil
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

				return &CharacterPlaytime{
					CharacterID:  characterID,
					FramesPlayed: framesPlayed,
					GameCount:    1,
					ReplayFiles:  []string{path},
				}
			}
			break
		}
	}
	return nil
}
