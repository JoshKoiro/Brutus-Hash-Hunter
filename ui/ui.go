package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/tj/go-spin"
)

func ShowSplash() {
	splashText, err := os.ReadFile("./config/splashscreen-manytools_org.txt")
	if err != nil {
		fmt.Println("Error getting splashscreen...\n", err)
		return
	}
	color.HiBlack(string(splashText))
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

func Spinner(text string) {
	s := spin.New()
	s.Set(spin.Spin1)
	for i := 0; i < 4; i++ {
		fmt.Printf("\r  \033[36m%v\033[m %s ", text, s.Next())
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Print("\033[2K") // Clear entire line
	fmt.Print("\r")      // Move cursor to the beginning of the line
}
