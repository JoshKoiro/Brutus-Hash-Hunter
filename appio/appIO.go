package appio

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadURL(fileName string, url string) {
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

func ReadFile(filePath string, userString string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
}

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

func GetLibrary() {
	//TODO: Function to read a JSON library file that contains the word lists that can be downloaded
}

func DeleteWordList(wordList string) {
	//TODO: function to delete the word list if the user requests this to be done when file closes
}
