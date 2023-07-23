package main

import (
	"brutus-hash-hunter/configFile"
	"brutus-hash-hunter/ui"
	"brutus-hash-hunter/workers"
	"fmt"
	"time"
)

var itemFound bool = false
var pItemFound *bool = &itemFound

type Result struct {
	StringVal string
	Match     bool
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
		*pItemFound = false

		for _, value := range configSettings.Wordlists {
			if *pItemFound {
				break
			}
			filePath = "./wordlists/" + value.Name + ".txt"
			fmt.Printf("\n\nSearching %v...\n", value.Name)
			workers.ProcessFile(filePath, userString, value.Name, pItemFound)

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
