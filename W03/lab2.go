package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"lab3/utils/crack"
)

func main() {
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	wordlistPath := "nord_vpn.txt"
	outputFilePath := "verbose_lab2.txt"

	// Create or truncate verbose output file
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Error creating verbose file: %v\n", err)
		return
	}
	defer outFile.Close()

	// Write to both console and file
	multiWriter := io.MultiWriter(os.Stdout, outFile)

	fmt.Fprintln(multiWriter, "==================================================")
	fmt.Fprintln(multiWriter, "          Lab #2: SHA-1 Password Cracker          ")
	fmt.Fprintln(multiWriter, "==================================================")
	fmt.Fprintf(multiWriter, "Target SHA-1 Hash : %s\n", targetHash)
	fmt.Fprintf(multiWriter, "Wordlist File     : %s\n\n", wordlistPath)

	// Open the wordlist file
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Fprintf(multiWriter, "Error opening wordlist: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	attempt := 0
	found := false

	for scanner.Scan() {
		attempt++
		password := strings.TrimRight(scanner.Text(), "\r\n")
		computedHash := crack.SHA1(password)

		if computedHash == targetHash {
			fmt.Fprintf(multiWriter, "[FOUND] Attempt #%d | Password: %s | Hash: %s\n", attempt, password, computedHash)
			fmt.Fprintln(multiWriter, "==================================================")
			fmt.Fprintf(multiWriter, "SUCCESS! Password found: '%s' on attempt %d\n", password, attempt)
			fmt.Fprintln(multiWriter, "==================================================")
			found = true
			break
		} else {
			fmt.Fprintf(multiWriter, "[TRY]   Attempt #%d | Password: %-20s | Hash: %s\n", attempt, password, computedHash)
		}
	}

	if !found {
		fmt.Fprintln(multiWriter, "\n[-] Password not found in the wordlist.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(multiWriter, "Error reading wordlist: %v\n", err)
	}

	fmt.Printf("\n[+] Verbose output saved to: %s\n", outputFilePath)
}

