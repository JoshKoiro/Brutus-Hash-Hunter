package main

import (
	"brutus-hash-hunter/appio"
	"brutus-hash-hunter/hashes"
	"brutus-hash-hunter/ui"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type Result struct {
	StringVal string
	Match     bool
}

func worker(workID int, userStringHash string, jobs <-chan string, results chan<- Result) {
	for n := range jobs {
		if userStringHash == hashes.SHA256(n) {
			//found a match
			results <- Result{StringVal: n, Match: true}
		} else {
			//no match
			results <- Result{StringVal: n, Match: false}
		}
	}
}

func compareHashesApp(filePath string, userString string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	lines := strings.Split(string(file), "\n")
	fileLength := len(lines)

	//creating channels
	jobs := make(chan string, fileLength+1)
	results := make(chan Result, fileLength+1)
	numCPU := runtime.NumCPU()
	startTime := time.Now()
	userStringHash := hashes.SHA256(userString)

	//Create workers
	for j := 1; j <= numCPU; j++ {
		go worker(j, userStringHash, jobs, results)
	}

	//fill up jobs
	for _, line := range lines {
		jobs <- line
	}

	for a := 1; a <= fileLength; a++ {
		result := <-results
		if result.Match {
			fmt.Printf("\nFound a match: %v on line %v \n", result.StringVal, a)
			break
		}
	}

	// Timer ends here
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

	//TODO: Check if wordlist is already downloaded - do not download if it is.

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
