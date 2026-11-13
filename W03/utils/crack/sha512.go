package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512 computes the SHA512 hash of a plain text string and returns the hex encoded string.
func SHA512(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

