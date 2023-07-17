package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func getAppMode() string {
	fmt.Println("Welcome to BrUTus. Please choose an option(1/2):")
	fmt.Println("1: check password against list")
	fmt.Println("2: check password hash against hashes")
	var programChoice string
	fmt.Scan(&programChoice)

	//TODO: validate the input

	return programChoice
}

func getFileName() string {
	var link string
	fmt.Println("Please select a wordlist:")
	fmt.Println("1: Xato-net 10 million passwords")
	//TODO: add more options here
	var wordlistChoice string
	fmt.Scan(&wordlistChoice)

	switch wordlistChoice {
	case "1":
		link = "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/xato-net-10-million-passwords.txt"
	//TODO: add more options here
	default:
		link = "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/xato-net-10-million-passwords.txt"
	}

	return link
}

func downloadLink(url string) {
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

func comparePassword(lineVal string, password string) bool {
	if lineVal == password {
		return true
	} else {
		return false
	}
}

func hashSHA256(value string) string {
	hash := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", hash)
}

func compareHash(lineVal string, password string) bool {
	var hashedLine string = hashSHA256(lineVal)
	var hashedPass string = hashSHA256(password)
	return hashedLine == hashedPass
}

func comparePasswordApp(file io.Reader, password string) {
	var iteration int = 1

	//timer starts here
	startTime := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if comparePassword(line, password) {
			fmt.Printf("password found at line %v!\n", iteration)
			break
		} else {
			iteration++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	//timer ends here
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	mDuration := duration.Milliseconds()
	fmt.Printf("Number of iterations: %v\n", iteration)
	fmt.Printf("Elapsed Time: %v ms\n", mDuration)
}

func compareHashesApp(file io.Reader, password string) {
	var iteration int = 1

	//timer starts here
	startTime := time.Now()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if compareHash(line, password) {
			fmt.Printf("password found at line %v!\n", iteration)
			break
		} else {
			iteration++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	//timer ends here
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	mDuration := duration.Milliseconds()
	fmt.Printf("Number of iterations: %v\n", iteration)
	fmt.Printf("Elapsed Time: %v ms\n", mDuration)
}

func main() {
	var appMode string
	var fileName string
	appMode = getAppMode()
	fileName = getFileName()

	downloadLink(fileName)

	for {
		filePath := "wordList.txt"

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		fmt.Println("\nPlease enter the password you wish to find:")
		var password string
		fmt.Scan(&password)

		switch appMode {
		case "1":
			comparePasswordApp(file, password)
		case "2":
			compareHashesApp(file, password)
		}

	}
}
