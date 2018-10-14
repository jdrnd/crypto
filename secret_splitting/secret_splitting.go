package secret_splitting

import (
	"encoding/hex"
	"math/rand"
)

// Takes in a plain text string, and number of ways to split
// Returns n HEX-encoded strings
func split_secret(secret string, n int) []string {
	split_secret_bytes := make([][]byte, 0, n)
	secret_buffer := []byte(secret)

	// Create n-1 random buffers and XOR the secret with each one
	for i := 0; i < n-1; i++ {
		iv := make([]byte, len(secret))
		_, _ = rand.Read(iv) // fill iv with a random initialization vector

		// XOR secret with the each IV
		for j := 0; j < len(secret); j++ {
			secret_buffer[j] ^= iv[j]
		}
		split_secret_bytes = append(split_secret_bytes, iv)
	}
	split_secret_bytes = append(split_secret_bytes, secret_buffer)

	split_secret := make([]string, 0, n)
	for _, buffer := range split_secret_bytes {
		split_secret = append(split_secret, hex.EncodeToString(buffer))
	}

	return split_secret
}

// Does not preform any message autentication so we have no idea if the correct
// number of split secrets were given or if the combining suceeds other than looking
// at the recombined secret
// TODO: impliment basic message authentication?
func combine_secret(split_secrets []string) string {
	len_secret := len(split_secrets[0]) / 2 // 2 hex characters per ascii character

	secret_bytes := make([]byte, len_secret)
	for i := 0; i < len(split_secrets); i++ {
		// Take advantage that go uninitialized values of bytes are zero
		secret_part_buffer, _ := hex.DecodeString(split_secrets[i])
		for j := 0; j < len_secret; j++ {
			secret_bytes[j] ^= secret_part_buffer[j]
		}
	}
	return string(secret_bytes)
}
