package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

// helper function to hash string with sha512
func SHA512(text string) string {
	b := []byte(text)
	h := sha512.Sum512(b)
	return hex.EncodeToString(h[:])
}
