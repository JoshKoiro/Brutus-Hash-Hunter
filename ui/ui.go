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

func SetWordList() string {
	var link string
	fmt.Println("Please select a wordlist:")
	fmt.Println("1: Xato-net 10 million passwords")
	//TODO: This should work with the appio.GetLibrary() function
	var wordlistChoice string
	fmt.Scan(&wordlistChoice)

	switch wordlistChoice {
	case "1":
		link = "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/xato-net-10-million-passwords.txt"
	//TODO: add more options here
	default:
		link = "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/xato-net-10-million-passwords.txt"
	}

	return link
}
