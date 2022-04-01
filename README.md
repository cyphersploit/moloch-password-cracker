# moloch-password-cracker
Password cracker written in Go

## About
Moloch password cracker is a simple password cracking tool written in Go. Moloch currently only supports SHA-256 unless manually changed.

## Usage
Moloch password cracker comes with a default wordlist of 370099 words, and a small hashed password list for testing purposes. Simply enter the name of the wordlist file, followed by the hashes file, including file extension.

Example usage:
go run main.go words_alpha.txt hashes.txt

## TODO
- Add support for multiple hashing algorithms
- Add support for cracking 1337 5p34k cracking
