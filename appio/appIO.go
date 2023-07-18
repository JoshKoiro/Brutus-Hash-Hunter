package appio

import (
	"io"
	"net/http"
	"os"
)

func OpenFile(fileName string) {

}

func ScanFile(fileName string) {

}

func DownloadURL(url string) {
	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Create local file
	out, err := os.Create("wordList.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Write body to file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		panic(err)
	}
}

func GetLibrary() {
	//TODO: Function to read a JSON library file that contains the word lists that can be downloaded
}

func DeleteWordList(wordList string) {
	//TODO: function to delete the word list if the user requests this to be done when file closes
}
