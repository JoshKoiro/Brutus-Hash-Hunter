package main

import (
	"brutus-hash-hunter/configFile"
	"brutus-hash-hunter/ui"
	"brutus-hash-hunter/workers"
	"time"

	"github.com/fatih/color"
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
			color.HiBlue("\n\nSearching %v...\n", value.Name)
			workers.ProcessFile(filePath, userString, value.Name, pItemFound)

			if !itemFound {
				color.Red("Item was not found in any of the wordlists")
			}
		}

		// Timer ends here
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		mDuration := duration.Milliseconds()
		color.Green("\n\nElapsed Time: %v ms\n\n", mDuration)

	}
}
