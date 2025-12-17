package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	var text = "rahasia"
	fmt.Printf("original : %s\n\n", text)

	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)

	_, _, _ = salt1, salt2, salt3
}
