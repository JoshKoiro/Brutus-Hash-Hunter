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

func processLines(lines []string, searchString string, resultChan chan<- bool, done chan struct{}) {
	hashedUserText := hashes.SHA256(searchString)
	for _, line := range lines {
		match := hashes.SHA256(line) == hashedUserText
		resultChan <- match

		// Check if any goroutine found a match
		select {
		case <-done:
			return
		default:
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

	numCPU := runtime.NumCPU()
	resultChan := make(chan bool)
	done := make(chan struct{})

	lineChunks := chunkLines(lines, numCPU)
	startTime := time.Now()

	for _, chunk := range lineChunks {
		go processLines(chunk, userString, resultChan, done)
	}

	// Check if any goroutine found a match
	for range lineChunks {
		match := <-resultChan
		if match {
			close(done) // Signal other goroutines to stop processing
			fmt.Println("Match found!")
			break
		} else {
			fmt.Println("No match found.")
		}
	}

	// Timer ends here
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	mDuration := duration.Milliseconds()
	fmt.Printf("Elapsed Time: %v ms\n", mDuration)
}

func chunkLines(lines []string, numChunks int) [][]string {
	chunkSize := (len(lines) + numChunks - 1) / numChunks // calculate chunk size
	var chunks [][]string
	for i := 0; i < len(lines); i += chunkSize {
		end := i + chunkSize
		if end > len(lines) {
			end = len(lines)
		}
		chunks = append(chunks, lines[i:end])
	}
	return chunks
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
