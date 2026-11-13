package crack

import (
	"crypto/md5"
	"encoding/hex"
)

// helper function to hash string with md5
func MD5(text string) string {
	b := []byte(text)
	h := md5.Sum(b)
	return hex.EncodeToString(h[:])
}
