package appio

import (
	"io"
	"net/http"
	"os"
)

func DownloadURL(fileName string, url string) {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	file, err := os.Create("./wordlists/" + fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		panic(err)
	}
}
