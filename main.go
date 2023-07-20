package main

import (
	"brutus-hash-hunter/appio"
	"brutus-hash-hunter/hashes"
	"brutus-hash-hunter/ui"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

func processLines(file *os.File, searchString string, resultChan chan<- bool) {
	hashedUserText := hashes.SHA256(searchString)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := hashes.SHA256(line) == hashedUserText
		resultChan <- match
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading file:", err)
	}
}

func compareHashesApp(filePath string, userString string) {

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	resultChan := make(chan bool)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	startTime := time.Now()
	for i := 0; i < numCPU; i++ {
		go processLines(file, userString, resultChan)
	}

	// Collect results
	for i := 0; i < numCPU; i++ {
		match := <-resultChan
		if match {
			fmt.Println("Match found!")
		} else {
			fmt.Println("No match found.")
		}
	}

	//timer ends here
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	mDuration := duration.Milliseconds()
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

		fmt.Println("\nPlease enter the text string you wish to compare:")
		var userString string
		fmt.Scan(&userString)

		switch progMode {
		case "1":
			// compareTextApp(file, userString)
		case "2":
			compareHashesApp(filePath, userString)
		}

	}
}
