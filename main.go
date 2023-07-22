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

func worker(workID int, userString string, jobs <-chan string, results chan<- Result) {
	for n := range jobs {
		if userString == hashes.SHA256(n) {
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

	//Create workers
	for j := 1; j <= numCPU; j++ {
		go worker(j, userString, jobs, results)
	}

	//fill up jobs
	for _, line := range lines {
		jobs <- line
	}

	for a := 1; a <= fileLength; a++ {
		result := <-results
		if result.Match {
			fmt.Printf("\nFound a match: %v \n", result.StringVal)
			fmt.Printf("Found on line: %v \n", a)
			fmt.Printf("Found in filename: %v \n", filePath)
			break
		}

		if a == fileLength {
			fmt.Println("No matches found")
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
	var wordList string
	wordList = ui.SetWordList()

	//TODO: Check if wordlist is already downloaded - do not download if it is.

	appio.DownloadURL(wordList, "Xato") //TODO: the second argument is the filename - have this be provided throguh the SetWordList function

	//TODO: give the ability to exit this loop and go back to redefine another progMode
	for {
		filePath := "Xato.txt" //TODO: the second argument is the filename - have this be provided throguh the SetWordList function

		fmt.Println("\nPlease enter the hash value you wish to find:")
		var userString string
		fmt.Scan(&userString)
		compareHashesApp(filePath, userString)

	}
}
