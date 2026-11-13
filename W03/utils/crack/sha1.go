package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 computes the SHA1 hash of a plain text string and returns the hex encoded string.
func SHA1(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

