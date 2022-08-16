package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	// <bin> -check <password> <bcrypt-hash>
	// <bin> -hash <password>
	// <bin> -hash <password> -check
	checkFlag := flag.Bool("check", false, "check a given password/bcrypt-hash match")
	hashFlag := flag.Bool("hash", false, "Bcrypt hash a given password")
	flag.Parse()

	fmt.Println(os.Args)

	password := os.Args[1]
	var hash string
	if *hashFlag {
		var err error
		if hash, err = HashPassword(password); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Password:", password)
		fmt.Println("Hash:    ", hash)
	}

	if *checkFlag {
		if hash == "" {
			hash = os.Args[2]
		}
		matches := CheckPasswordHash(password, hash)
		fmt.Println("Match:   ", matches)
	}

}
