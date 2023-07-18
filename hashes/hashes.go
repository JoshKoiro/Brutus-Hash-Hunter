package hashes

import (
	"crypto/sha256"
	"fmt"
)

func HashSHA256(value string) string {
	hash := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", hash)
}

func CompareText(lineVal string, password string) bool {
	if lineVal == password {
		return true
	} else {
		return false
	}
}
