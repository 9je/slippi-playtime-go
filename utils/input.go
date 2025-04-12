package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sqweek/dialog"
)

func GetSlippiCode() string {
	var slippiCode string
	fmt.Print("Enter your Slippi code or name: ")
	fmt.Scanln(&slippiCode)
	return slippiCode
}

func PromptForDirectory() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %v", err)
	}
	defaultDir := filepath.Join(homeDir, "Documents", "Slippi")
	dir, err := dialog.Directory().Title("Select Your Slippi Replay Folder").SetStartDir(defaultDir).Browse()
	if err != nil {
		log.Fatalf("Error selecting folder: %v", err)
	}
	return dir
}
