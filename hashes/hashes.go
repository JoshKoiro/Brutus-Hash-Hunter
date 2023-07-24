package hashes

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func SHA256(value string) string {
	hash := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", hash)
}

func SHA512(value string) string {
	hash := sha512.Sum512([]byte(value))
	return fmt.Sprintf("%x", hash)
}

func MD5(value string) string {
	inputBytes := []byte(value)
	hash := md5.Sum(inputBytes)
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
