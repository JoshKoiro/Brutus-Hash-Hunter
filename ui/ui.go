package ui

import (
	"brutus-hash-hunter/appio"
	"fmt"
)

func ShowSplash() {
	splashText, err := appio.ReadFileAsString("splashscreen-manytools_org.txt")
	if err != nil {
		fmt.Println("Error getting splashscreen...\n", err)
		return
	}

	fmt.Println(splashText)
}

func FilesMessage(wordListsLength int) {
	fmt.Printf("\nNumber of files loaded from config.yaml: %v\n", wordListsLength)
	fmt.Printf("\nFiles loaded:\n")
}

func AskForHash() string {
	fmt.Println("\nPlease enter the hash value you wish to find:")
	var hash string
	fmt.Scan(&hash)
	return hash
}
