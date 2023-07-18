package ui

import (
	"brutus-hash-hunter/appio"
	"fmt"
)

func ShowSplash() {
	//show splashscreen text file on startup
	splashText, err := appio.ReadFileAsString("splashscreen-manytools_org.txt")
	if err != nil {
		fmt.Println("Error getting splashscreen...\n", err)
		return
	}

	fmt.Println(splashText)
}

func SetAppMode() string {
	fmt.Println("Welcome to Brutus! Please choose an option(1/2):")
	fmt.Println("1: check text against list")
	fmt.Println("2: check text hash against hashes")
	var progMode string
	fmt.Scan(&progMode)

	//TODO: validate the input

	return progMode
}

func SetWordList() string {
	var link string
	fmt.Println("Please select a wordlist:")
	fmt.Println("1: Xato-net 10 million passwords")
	//TODO: add more options here
	//TODO: This should probably be broken out into a JSON or .yaml file.
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
