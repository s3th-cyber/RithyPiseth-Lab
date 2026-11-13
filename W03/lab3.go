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
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	wordlistPath := "nord_vpn.txt"
	outputFilePath := "verbose_lab3.txt"

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
	fmt.Fprintln(multiWriter, "         Lab #3: SHA-512 Password Cracker         ")
	fmt.Fprintln(multiWriter, "==================================================")
	fmt.Fprintf(multiWriter, "Target SHA-512 Hash : %s\n", targetHash)
	fmt.Fprintf(multiWriter, "Wordlist File       : %s\n\n", wordlistPath)

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
		computedHash := crack.SHA512(password)

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

