package ui

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ShowSplash() {
	splashText, err := os.ReadFile("./config/splashscreen-manytools_org.txt")
	if err != nil {
		fmt.Println("Error getting splashscreen...\n", err)
		return
	}
	color.Cyan(string(splashText))
}

func WordlistInfo(wordListsLength int) {
	color.Green("\nNumber of files loaded from config.yaml: %v\n", wordListsLength)
	fmt.Printf("\nFiles loaded:\n")
}

func AskForHash() string {
	color.Yellow("\nPlease enter the SHA256 hash value you wish to find (press ctrl-c to exit):")
	var hash string
	fmt.Scan(&hash)
	return hash
}
