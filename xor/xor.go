package crypto

import (
	"io/ioutil"

	"github.com/gonum/stat"
)

const MAX_KEY_LEN = 64
const OUTLIER_DEVIATION = 1.5
const BASIC_ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ '"

// Returns a slice of bytes
func encrypt_string(plaintext string, key string) []byte {
	// Note: this does NOT clear the plaintext string
	data := []byte(plaintext)

	keypos := 0

	for datapos := 0; datapos < len(data); datapos++ {
		keypos = keypos % len(key)

		data[datapos] ^= key[keypos]

		keypos++
	}
	return data
}

func decrypt_string(plaintext string, key string) []byte {
	return encrypt_string(plaintext, key)
}

func encrypt_file(filename string, key string) []byte {
	file_contents, _ := ioutil.ReadFile(filename)
	return encrypt_string(string(file_contents), key)
}

func get_key_length(cyphertext []byte) int {
	// Assume MAX_KEY_LEN is very small compared to the message length
	num_matches := make(map[int]float32)
	for step := 1; step <= MAX_KEY_LEN; step++ {
		matches := 0
		total := 0
		for i := 0; i < len(cyphertext); i++ {
			for j := i + step; j < len(cyphertext); j += step {
				total++
				if cyphertext[i] == cyphertext[j] {
					matches++
				}
			}
		}
		num_matches[step] = float32(matches) / float32(total)
	}

	matches_only := make([]float64, 0, MAX_KEY_LEN)
	for _, value := range num_matches {
		matches_only = append(matches_only, float64(value))
	}

	// Apply a z-test to find outliers
	mean, stddev := stat.MeanStdDev(matches_only, nil)

	lowest_outlier := 0
	for i := 1; i <= MAX_KEY_LEN; i++ {
		if (float64(num_matches[i])-mean)/stddev >= OUTLIER_DEVIATION {
			lowest_outlier = i
			break
		}
	}

	return lowest_outlier
}

func break_cypher(cyphertext string) string {
	key_length := get_key_length([]byte(cyphertext))
	key := get_multi_character_xor_key([]byte(cyphertext), key_length)

	return string(decrypt_string(cyphertext, string(key)))
}

func get_multi_character_xor_key(cyphertext []byte, key_length int) []byte {
	// Could paralellize this if required
	key := make([]byte, key_length)
	for offset := 0; offset < key_length; offset++ {
		bytes_for_key_char := make([]byte, 0) // stores all bytes xor'd with speicific byte of key
		index := offset
		for index < len(cyphertext) {
			bytes_for_key_char = append(bytes_for_key_char, cyphertext[index])
			index += key_length
		}
		key[offset] = get_single_character_xor_key(bytes_for_key_char)
	}
	return key
}

func get_single_character_xor_key(data []byte) byte {
	// simple niave approach here, we just assume the most common letter is the space (ASCII encoding)
	letter_counts := make(map[byte]int)
	for _, letter := range data {
		letter_counts[letter]++
	}
	highest_count := 0
	most_common_letter := byte(0)
	for letter, count := range letter_counts {
		if count > highest_count {
			highest_count = count
			most_common_letter = letter
		}
	}

	return (most_common_letter ^ byte(' '))
}
