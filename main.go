package main

import (
	"brutus-hash-hunter/appio"
	"brutus-hash-hunter/hashes"
	"brutus-hash-hunter/ui"
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var itemFound bool = false

type Result struct {
	StringVal string
	Match     bool
}

// type FileLine struct {
// 	lineNbr int
// 	text    string
// }

func worker(workID int, userString string, jobs <-chan string, results chan<- Result) {
	for n := range jobs {
		if userString == hashes.SHA256(n) {
			//found a match
			results <- Result{StringVal: n, Match: true}
		}
	}
}

func ProcessFiles(filePath string, userString string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := strings.Split(string(file), "\n")
	fileLength := len(lines)

	//creating channels
	jobs := make(chan string, fileLength+1)
	results := make(chan Result)
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

	fmt.Printf("length of jobs: %v\n", len(jobs))

	// Timer ends here
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	mDuration := duration.Milliseconds()
	fmt.Printf("Elapsed Time: %v ms\n", mDuration)

	for result := range results {
		if result.Match {
			fmt.Printf("\nFound a match: %v \n", result.StringVal)
			fmt.Printf("Found in filename: %v \n", filePath)
		} else {
			fmt.Println("No matches found")
			break
		}
	}
}

func readFile(filePath string, userString string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 1
	iterate := &i
	for scanner.Scan() {
		if itemFound {
			wg.Wait()
			return
		}
		wg.Add(1)
		go simpleWorker(userString, *iterate, scanner.Text())
		*iterate = *iterate + 1
	}
	wg.Wait()
}

func simpleWorker(userString string, lineNbr int, line string) {
	if userString == hashes.SHA256(line) {
		//found a match
		fmt.Printf("\n Found a match! %v", line)
		fmt.Printf("\n in file: ")
		fmt.Printf("\n on line: %v", lineNbr)
		itemFound = true
	}
	wg.Done()
}

func main() {
	ui.ShowSplash()
	var wordList string = ui.SetWordList()

	//TODO: Check if wordlist is already downloaded - do not download if it is.

	appio.DownloadURL(wordList, "Xato") //TODO: the second argument is the filename - have this be provided throguh the SetWordList function

	//TODO: give the ability to exit this loop and go back to redefine another progMode
	for {
		filePath := "Xato.txt" //TODO: the second argument is the filename - have this be provided throguh the SetWordList function

		fmt.Println("\nPlease enter the hash value you wish to find:")
		var userString string
		fmt.Scan(&userString)
		// processFiles(filePath, userString)
		startTime := time.Now()
		itemFound = false
		readFile(filePath, userString)
		// Timer ends here
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		mDuration := duration.Milliseconds()
		fmt.Printf("\nElapsed Time: %v ms", mDuration)

	}
}
