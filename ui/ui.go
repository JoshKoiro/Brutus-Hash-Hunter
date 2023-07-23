package ui

import (
	"brutus-hash-hunter/appio"
	"fmt"

	"github.com/fatih/color"
)

func ShowSplash() {
	splashText, err := appio.ReadFileAsString("splashscreen-manytools_org.txt")
	if err != nil {
		fmt.Println("Error getting splashscreen...\n", err)
		return
	}
	color.Cyan(splashText)
	// fmt.Println(splashText)
}

func FilesMessage(wordListsLength int) {
	color.Green("\nNumber of files loaded from config.yaml: %v\n", wordListsLength)
	fmt.Printf("\nFiles loaded:\n")
}

func AskForHash() string {
	color.Yellow("\nPlease enter the SHA256 hash value you wish to find:")
	var hash string
	fmt.Scan(&hash)
	return hash
}
