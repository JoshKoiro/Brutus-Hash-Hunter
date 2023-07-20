package main

import (
	"brutus-hash-hunter/appio"
	"brutus-hash-hunter/compare"
	"brutus-hash-hunter/ui"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func compareTextApp(file io.Reader, userString string) {
	var iteration int = 1

	//timer starts here
	startTime := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if compare.CompareText(line, userString) {
			fmt.Printf("userString found at line %v!\n", iteration)
			break
		} else {
			iteration++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	//timer ends here
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	mDuration := duration.Milliseconds()
	fmt.Printf("Number of iterations: %v\n", iteration)
	fmt.Printf("Elapsed Time: %v ms\n", mDuration)
}

func compareHashesApp(file io.Reader, userString string) {
	var iteration int = 1

	//timer starts here
	startTime := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if compare.CompareSHA256(line, userString) {
			fmt.Printf("userString found at line %v!\n", iteration)
			break
		} else {
			iteration++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	//timer ends here
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	mDuration := duration.Milliseconds()
	fmt.Printf("Number of iterations: %v\n", iteration)
	fmt.Printf("Elapsed Time: %v ms\n", mDuration)
}

func main() {
	ui.ShowSplash()
	var progMode string
	var wordList string
	progMode = ui.SetAppMode()
	wordList = ui.SetWordList()

	appio.DownloadURL(wordList, "Xato") //TODO: the second argument is the filename - have this be provided throguh the SetWordList function

	//TODO: give the ability to exit this loop and go back to redefine another progMode
	for {
		filePath := "Xato.txt" //TODO: the second argument is the filename - have this be provided throguh the SetWordList function

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		fmt.Println("\nPlease enter the text string you wish to compare:")
		var userString string
		fmt.Scan(&userString)

		switch progMode {
		case "1":
			compareTextApp(file, userString)
		case "2":
			compareHashesApp(file, userString)
		}

	}
}
