package compare

import "brutus-hash-hunter/hashes"

func CompareSHA256(lineVal string, password string) bool {
	var hashedLine string = hashes.HashSHA256(lineVal)
	var hashedPass string = hashes.HashSHA256(password)
	return hashedLine == hashedPass
}
