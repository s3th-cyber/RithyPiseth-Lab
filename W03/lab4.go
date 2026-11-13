package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
)

func main() {
	// Lab 4: Mini-CTF (Cat Leak)
	// From the url example:
	// p1crypt=b6ccb4ece5454dca and p1crypt=e51778b3e239ebc2
	// When we put them together, it becomes a 32-character MD5 hash:
	leakedHash := "b6ccb4ece5454dca" + "e51778b3e239ebc2"

	fmt.Println("Target Leaked Hash:", leakedHash)

	// Note from prompt:
	// "I love cats and my password is something related to cats"
	// Also regex given: ^cryptoCTF{(?:\x6d\x65\x6f\x77){2}}$
	// \x6d = 'm', \x65 = 'e', \x6f = 'o', \x77 = 'w' => "meow"
	// repeated 2 times => "meowmeow"

	word := "meowmeow"

	// compute MD5 of candidate
	b := []byte(word)
	h := md5.Sum(b)
	computedHash := hex.EncodeToString(h[:])

	fmt.Println("Testing word      :", word)
	fmt.Println("Computed Hash     :", computedHash)

	// check if hash matches leaked hash
	if computedHash == leakedHash {
		fmt.Println("\n=> Match found!")
		flag := "cryptoCTF{" + word + "}"
		fmt.Println("Flag:", flag)

		// double check with the lab regex
		pattern := `^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`
		matched, _ := regexp.MatchString(pattern, flag)
		fmt.Println("Regex check passed:", matched)
	}
}
