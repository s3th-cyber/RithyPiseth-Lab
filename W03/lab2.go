package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab3/utils/crack"
)

func main() {
	// target SHA1 hash to crack from lab 2
	target := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	// open wordlist file
	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error: cannot open wordlist file!")
		return
	}
	defer file.Close()

	// save verbose output to text file
	outFile, err := os.Create("verbose_lab2.txt")
	if err != nil {
		fmt.Println("Error creating output file!")
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	fmt.Println("Starting to crack SHA1 hash...")

	// read line by line using loop
	for scanner.Scan() {
		count++
		password := strings.TrimSpace(scanner.Text())

		// hash with sha1 using helper
		hash := crack.SHA1(password)

		// print and log verbose
		log := fmt.Sprintf("[%d] Testing: %s => %s\n", count, password, hash)
		fmt.Print(log)
		outFile.WriteString(log)

		// compare hash
		if hash == target {
			result := fmt.Sprintf("\nPassword Found!\nPassword: %s\nAttempts: %d\n", password, count)
			fmt.Print(result)
			outFile.WriteString(result)
			break
		}
	}
}
