package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"time"
)

var link string

func askRole() {
	fmt.Println("Welcome to BrUTus. Please choose an option(1/2):")
	fmt.Println("1: check password against list")
	fmt.Println("2: check password hash against hashes")
	var programChoice string
	fmt.Scan(&programChoice)

	fmt.Println("Please select a wordlist:")
	fmt.Println("1: Xato-net 10 million passwords")
	//TODO: will add more options here
	var wordlistChoice string
	fmt.Scan(&wordlistChoice)

	switch wordlistChoice {
	case "1":
		link = "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/xato-net-10-million-passwords.txt"
	//TODO: will add more options here
	default:
		link = "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/xato-net-10-million-passwords.txt"
	}

	getLink(link)

	if programChoice == "y" {
		var lineVal string
		var password string
		comparePassword(lineVal, password)
		//user has selected to compare password against list
	} else {
		var lineVal string
		var password string
		compareHash(lineVal, password)
		//user has selected to compare hashes
	}
}

func getLink(url string) string {
	return url
}

func comparePassword(lineVal string, password string) bool {
	return lineVal == password
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

func main() {
	askRole()
	for {
		filePath := "C:/Users/Josh/Desktop/learning-go/SecLists/Passwords/xato-net-10-million-passwords.txt"

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		//ask the user what password to look for
		fmt.Println("\nPlease enter the password you wish to find:")
		var password string
		fmt.Scan(&password)
		var iteration int = 1
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

		endTime := time.Now()
		duration := endTime.Sub(startTime)
		mDuration := duration.Milliseconds()
		fmt.Printf("Number of iterations: %v\n", iteration)
		fmt.Printf("Elapsed Time: %v ms\n", mDuration)
	}
}
