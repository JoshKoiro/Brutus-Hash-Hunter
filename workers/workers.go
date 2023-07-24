package workers

import (
	"brutus-hash-hunter/hashes"
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
)

var wg sync.WaitGroup

func Deploy(filePath string, userString string, fileName string, itemFound *bool) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()
	fileLine := bufio.NewScanner(file)
	i := 1
	iterate := &i
	for fileLine.Scan() {
		if *itemFound {
			wg.Wait()
			return
		}
		wg.Add(1)
		go Worker(userString, *iterate, fileLine.Text(), fileName, itemFound)
		*iterate = *iterate + 1
	}
	wg.Wait()
	color.HiBlack("searched %v lines...", *iterate)
}

func Worker(userString string, lineNbr int, line string, fileName string, itemFound *bool) {
	if userString == hashes.SHA256(line) {
		boldGreen := color.New(color.FgGreen).Add(color.Bold).Add(color.Underline)
		fmt.Println("\n=========================================")
		boldGreen.Printf("\nFound a match!")
		color.HiGreen(" %v", line)
		color.HiBlue("\non line %v", lineNbr)
		color.HiBlue("in %v\n", fileName)
		fmt.Println("\n=========================================")
		*itemFound = true
	}
	wg.Done()
}
