package workers

import (
	"brutus-hash-hunter/hashes"
	"bufio"
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func ProcessFile(filePath string, userString string, fileName string, itemFound *bool) {
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
		if *itemFound {
			wg.Wait()
			return
		}
		wg.Add(1)
		go Worker(userString, *iterate, scanner.Text(), fileName, itemFound)
		*iterate = *iterate + 1
	}
	wg.Wait()
	fmt.Printf("while searching %v lines...", *iterate)
}

func Worker(userString string, lineNbr int, line string, fileName string, itemFound *bool) {
	if userString == hashes.SHA256(line) {
		//found a match
		fmt.Printf("\n Found a match! %v", line)
		fmt.Printf("\n on line %v", lineNbr)
		fmt.Printf(" in %v\n", fileName)
		*itemFound = true
	}
	wg.Done()
}
