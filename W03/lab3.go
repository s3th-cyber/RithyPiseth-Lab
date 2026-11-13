package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab3/utils/crack"
)

func main() {
	// target SHA512 hash to crack from lab 3
	target := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	// open the wordlist file
	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error: cannot open wordlist file!")
		return
	}
	defer file.Close()

	// output file for verbose
	outFile, err := os.Create("verbose_lab3.txt")
	if err != nil {
		fmt.Println("Error creating output file!")
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	fmt.Println("Starting to crack SHA512 hash...")

	// loop through each password in wordlist
	for scanner.Scan() {
		count++
		password := strings.TrimSpace(scanner.Text())

		// hash with sha512 helper
		hash := crack.SHA512(password)

		// write verbose to terminal and file
		log := fmt.Sprintf("[%d] Testing: %s => %s\n", count, password, hash)
		fmt.Print(log)
		outFile.WriteString(log)

		// check match
		if hash == target {
			result := fmt.Sprintf("\nPassword Found!\nPassword: %s\nAttempts: %d\n", password, count)
			fmt.Print(result)
			outFile.WriteString(result)
			break
		}
	}
}
