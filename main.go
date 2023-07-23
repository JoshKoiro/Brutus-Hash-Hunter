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

func readFile(filePath string, userString string, fileName string) {
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
		go simpleWorker(userString, *iterate, scanner.Text(), fileName)
		*iterate = *iterate + 1
	}
	wg.Wait()
	fmt.Printf("while searching %v lines...", *iterate)
}

func simpleWorker(userString string, lineNbr int, line string, fileName string) {
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

	fmt.Printf("\nNumber of files loaded from config: %v\n", len(configSettings.Wordlists))
	fmt.Printf("\nFiles loaded:\n")

	var index int = 1
	for _, value := range configSettings.Wordlists {
		if _, err := os.Stat("./wordlists/" + value.Name + ".txt"); err == nil {
			fmt.Printf("(%v/%v) Already downloaded %v\n", index, len(configSettings.Wordlists), value.Name)
		} else {
			fmt.Printf("(%v/%v) Downloading %v...\n", index, len(configSettings.Wordlists), value.Name)
			appio.DownloadURL(value.Name, value.Link)
		}
		index++
	}

	//TODO: give the ability to exit this loop and go back to redefine another progMode
	for {
		var filePath string

		fmt.Println("\nPlease enter the hash value you wish to find:")
		var userString string
		fmt.Scan(&userString)
		startTime := time.Now()
		itemFound = false
		for _, value := range configSettings.Wordlists {
			if itemFound {
				break
			}
			filePath = "./wordlists/" + value.Name + ".txt"
			fmt.Printf("\n\nSearching %v...\n", value.Name)
			readFile(filePath, userString, value.Name)

			if !itemFound {
				fmt.Printf("Item was not found in any of the wordlists")
			}
		}
		// Timer ends here
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		mDuration := duration.Milliseconds()
		fmt.Printf("\n\nElapsed Time: %v ms\n\n", mDuration)

	}
}
