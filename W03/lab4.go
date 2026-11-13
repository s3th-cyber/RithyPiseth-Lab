package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("          Lab #4: Mini-CTF (Cat Leak)             ")
	fmt.Println("==================================================")

	// Step 1: Decode the leaked Google search query
	encodedQuery := "%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32"
	decodedQuery, err := url.QueryUnescape(encodedQuery)
	if err != nil {
		fmt.Printf("Error decoding URL: %v\n", err)
		return
	}
	fmt.Printf("[Step 1] Decoded Search Query (Leaked Hash):\n         %s\n\n", decodedQuery)

	// Step 2: Show the URL split hashes also match
	p1Part1 := "b6ccb4ece5454dca"
	p1Part2 := "e51778b3e239ebc2"
	combinedHash := p1Part1 + p1Part2
	fmt.Printf("[Step 2] Combined URL Parameters (p1crypt + p1crypt):\n         %s + %s = %s\n", p1Part1, p1Part2, combinedHash)
	fmt.Printf("         Hash Length: %d chars (MD5 Hash)\n\n", len(combinedHash))

	// Step 3: Crack the cat password
	// The user mentions "I love cats and my password is something related to cats"
	catWordCandidates := []string{
		"cat", "cats", "kitty", "kitten", "meow", "meowmeow", "purr", "purrpurr",
	}

	var crackedPassword string
	for _, word := range catWordCandidates {
		h := md5.Sum([]byte(word))
		hexHash := hex.EncodeToString(h[:])
		if hexHash == combinedHash {
			crackedPassword = word
			fmt.Printf("[Step 3] Cracking Cat Password:\n")
			fmt.Printf("         [MATCH] Password: '%s' -> MD5: %s\n\n", word, hexHash)
			break
		}
	}

	if crackedPassword == "" {
		fmt.Println("[-] Password could not be matched with candidates.")
		return
	}

	// Step 4: Construct the flag
	flag := fmt.Sprintf("cryptoCTF{%s}", crackedPassword)
	fmt.Printf("[Step 4] Constructed Flag: %s\n\n", flag)

	// Step 5: Verify with the regex provided in the lab:
	// Verify answer (Regex): ^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$
	// Note: \x6d='m', \x65='e', \x6f='o', \x77='w' -> "meow" repeated twice = "meowmeow"
	verifyRegex := `^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`
	matched, err := regexp.MatchString(verifyRegex, flag)
	if err != nil {
		fmt.Printf("Regex error: %v\n", err)
		return
	}

	fmt.Println("==================================================")
	fmt.Printf("Verification Regex : %s\n", verifyRegex)
	fmt.Printf("Regex Matches Flag : %v\n", matched)
	fmt.Println("==================================================")
	if matched {
		fmt.Printf("FINAL FLAG: %s\n", flag)
	}
	fmt.Println("==================================================")
}

