package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

// helper function to hash string with sha1
func SHA1(text string) string {
	b := []byte(text)
	h := sha1.Sum(b)
	return hex.EncodeToString(h[:])
}
