package crack

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 computes the MD5 hash of a plain text string and returns the hex encoded string.
func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

