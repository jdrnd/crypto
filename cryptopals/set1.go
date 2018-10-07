package cryptopals

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func DecryptECB(cyphertext []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key) // runs in 128-bit mode
	position := 0

	for (position + 16) <= len(cyphertext) {
		block.Decrypt(cyphertext[position:position+16], cyphertext[position:position+16])
		position += 16
	}
	return cyphertext
}

/*
Challenges 1-6 were previous implimented in another module
*/
func set1Challenge7(cyphertext_file string) string {
	file_contents, _ := ioutil.ReadFile(cyphertext_file)
	encrypted_data, _ := base64.StdEncoding.DecodeString(string(file_contents))

	key := []byte("YELLOW SUBMARINE")
	plaintext := DecryptECB([]byte(encrypted_data), key)
	return string(plaintext)
}

func set1Challenge8() {
	/*
		For each input, construct a set containing each encrypted block. If a new encrypted block is already in the set the cypher must be ECB with a repeated input block
		(We can do this since the output space is 2^128 so the probablility of collisions is negligible)
	*/
	fmt.Println("IMPLIMENT ME PLEASE!!")
}
