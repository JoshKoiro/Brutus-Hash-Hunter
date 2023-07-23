package main

import (
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

func hashFile(filePath string, userString string, fileName string) {
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
		go Worker(userString, *iterate, scanner.Text(), fileName)
		*iterate = *iterate + 1
	}
	wg.Wait()
	fmt.Printf("while searching %v lines...", *iterate)
}

func Worker(userString string, lineNbr int, line string, fileName string) {
	if userString == hashes.SHA256(line) {
		//found a match
		fmt.Printf("\n Found a match! %v", line)
		fmt.Printf("\n on line %v", lineNbr)
		fmt.Printf(" in %v\n", fileName)
		itemFound = true
	}
	wg.Done()
}

func main() {
	ui.ShowSplash()
	configSettings := new(configFile.Config)
	configSettings.Initialize()

	var wordListLength int = len(configSettings.Wordlists)

	ui.FilesMessage(wordListLength)
	configSettings.DownloadFiles()

	for {
		var filePath string
		var userString string
		pUserString := &userString
		*pUserString = ui.AskForHash()

		startTime := time.Now()
		itemFound = false

		for _, value := range configSettings.Wordlists {
			if itemFound {
				break
			}
			filePath = "./wordlists/" + value.Name + ".txt"
			fmt.Printf("\n\nSearching %v...\n", value.Name)
			hashFile(filePath, userString, value.Name)

			if !itemFound {
				fmt.Printf("Item was not found in any of the wordlists")
			}
		}

		// Timer ends here
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		mDuration := duration.Milliseconds()
		fmt.Printf("\n\nElapsed Time: %v seconds\n\n", mDuration)

	}
}
