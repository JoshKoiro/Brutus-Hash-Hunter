package main

import (
	"brutus-hash-hunter/appio"
	"brutus-hash-hunter/configFile"
	"brutus-hash-hunter/hashes"
	"brutus-hash-hunter/ui"
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var itemFound bool = false

type Result struct {
	StringVal string
	Match     bool
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
	configSettings := new(configFile.Config)
	configSettings.Initialize()

	fmt.Printf("\nNumber of files loaded from config: %v\n", len(configSettings.Wordlists))
	fmt.Printf("\nFiles loaded:\n")

	var index int = 1
	for _, value := range configSettings.Wordlists {
		if _, err := os.Stat(value.Name + ".txt"); err == nil {
			fmt.Printf("(%v/%v) Already downloaded %v\n", index, len(configSettings.Wordlists), value.Name)
		} else {
			fmt.Printf("(%v/%v) Downloading %v...\n", index, len(configSettings.Wordlists), value.Name)
			appio.DownloadURL(value.Name, value.Link)
		}
		index++
	}

	//TODO: give the ability to exit this loop and go back to redefine another progMode
	for {
		filePath := "Xato-net-10-million-passwords.txt" //TODO: the second argument is the filename - have this be provided throguh the SetWordList function

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
