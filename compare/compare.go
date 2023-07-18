package compare

import "brutus-hash-hunter/hashes"

func CompareText(lineVal string, password string) bool {
	return lineVal == password
}

func CompareSHA256(lineVal string, password string) bool {
	var hashedLine string = hashes.SHA256(lineVal)
	var hashedPass string = hashes.SHA256(password)
	return hashedLine == hashedPass
}

func CompareMD5(lineVal string, password string) bool {
	var hashedLine string = hashes.MD5(lineVal)
	var hashedPass string = hashes.MD5(password)
	return hashedLine == hashedPass
}
