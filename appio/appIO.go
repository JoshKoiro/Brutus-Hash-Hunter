package appio

import (
	"io"
	"net/http"
	"os"
)

func ReadFileAsString(filename string) (string, error) {
	// Read the file contents
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	str := string(content)

	return str, nil
}

func ScanFile(fileName string) {

}

func DownloadURL(url string, fileName string) {
	// Send GET request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Create local file
	out, err := os.Create(fileName + ".txt")
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
