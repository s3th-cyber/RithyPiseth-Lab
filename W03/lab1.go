package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab3/utils/crack"
)

func main() {
	// hash that we want to crack 
	target := "6a85dfd77d9cb35770c9dc6728d73d3f"

	// open the wordlist file
	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error: cannot open wordlist file!")
		return
	}
	defer file.Close()

	// create file to save output (verbose.txt)
	outFile, err := os.Create("verbose_lab1.txt")
	if err != nil {
		fmt.Println("Error creating output file!")
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	fmt.Println("Starting to crack MD5 hash...")

	// loop every password in wordlist
	for scanner.Scan() {
		count++
		// remove newline and spaces
		password := strings.TrimSpace(scanner.Text())

		// compute hash using helper in utils/crack/md5.go
		hash := crack.MD5(password)

		// print attempt and write to file
		log := fmt.Sprintf("[%d] Testing: %s => %s\n", count, password, hash)
		fmt.Print(log)
		outFile.WriteString(log)

		// check if match with target hash
		if hash == target {
			result := fmt.Sprintf("\nPassword Found!\nPassword: %s\nAttempts: %d\n", password, count)
			fmt.Print(result)
			outFile.WriteString(result)
			break
		}
	}
}
