package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"
)

// TODO add support for multiple hashes - for loop through array of hashes
// TODO add L337 SPE4K checking

func main() {
	// get command line argument (txt file name)
	wordlist := os.Args[1]
	hashes := os.Args[2]
	readFile(wordlist, hashes)
}

func readFile(wordlist string, hasheslist string) {

	// open hasheslist
	hashesFile, err := os.Open(hasheslist)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	hashesScanner := bufio.NewScanner(hashesFile)
	hashesScanner.Split(bufio.ScanLines)
	var hasheslines []string

	for hashesScanner.Scan() {
		hasheslines = append(hasheslines, hashesScanner.Text())
	}

	hashesFile.Close()

	// open wordlist
	wordFile, err := os.Open(wordlist)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	wordScanner := bufio.NewScanner(wordFile)
	wordScanner.Split(bufio.ScanLines)
	var txtlines []string

	for wordScanner.Scan() {
		txtlines = append(txtlines, wordScanner.Text())
	}

	wordFile.Close()

	// Work on words in dictionary
	for _, password := range hasheslines {
		for _, wordLline := range txtlines {

			// hashed lowecase word
			data := []byte(strings.TrimSpace(wordLline))
			hash := sha256.Sum256(data)

			// hashed capitalized word
			capitalData := []byte(strings.TrimSpace(strings.Title(wordLline)))
			capitalHash := sha256.Sum256(capitalData)

			// final hashes to string
			hashed := fmt.Sprintf("%x", hash)
			capitalHashed := fmt.Sprintf("%x", capitalHash)

			if password == hashed {
				// check lowercase
				saveFound(wordLline, hashed)
			} else if password == capitalHashed {
				// check capitalized
				saveFound(strings.Title(wordLline), hashed)
			}
		}
	}
}

// create text file if doesnt exist, or append to if it does
func saveFound(password string, hash string) {
	passwordData := []byte(password)
	hashData := []byte(hash)

	file, err := os.OpenFile("recovered.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := fmt.Sprintf("--------------------\nPassword: %s\nHash: %s\n--------------------", passwordData, hashData)
	_, err = fmt.Fprintln(file, newLine)
	if err != nil {
		fmt.Println(err)
	}
}
