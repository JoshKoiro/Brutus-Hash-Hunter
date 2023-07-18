package main

import (
	"brutus-hash-hunter/appio"
	"brutus-hash-hunter/compare"
	"brutus-hash-hunter/hashes"
	"brutus-hash-hunter/ui"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func compareTextApp(file io.Reader, password string) {
	var iteration int = 1

	//timer starts here
	startTime := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if hashes.CompareText(line, password) {
			fmt.Printf("password found at line %v!\n", iteration)
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

func compareHashesApp(file io.Reader, password string) {
	var iteration int = 1

	//timer starts here
	startTime := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if compare.CompareSHA256(line, password) {
			fmt.Printf("password found at line %v!\n", iteration)
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
	var progMode string
	var wordList string
	progMode = ui.SetAppMode()
	wordList = ui.SetWordList()

	//downloads the selected textlist for quick iteration
	appio.DownloadURL(wordList)

	//main application loop
	//TODO: give the ability to exit this loop and go back to redefine another progMode
	for {
		filePath := "wordList.txt"

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		//get the requested password or hash from the user
		fmt.Println("\nPlease enter the password or hash you wish to compare:")
		var password string
		fmt.Scan(&password)

		switch progMode {
		case "1":
			compareTextApp(file, password)
		case "2":
			compareHashesApp(file, password)
		}

	}
}
