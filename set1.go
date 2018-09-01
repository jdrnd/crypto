package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Challenge 1

func buffer_to_b64_string(buffer []byte) string {
	return base64.StdEncoding.EncodeToString(buffer)
}

func challenge1() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	desiredOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	buffer := hex_string_to_bytes(input)
	passed := (buffer_to_b64_string(buffer) == desiredOutput)
	if passed {
		fmt.Println("Challenge 1 Passed")
	} else {
		fmt.Println("Challenge 1 Failed")
	}
}

// Challenge 2

func xorTogether(buffer1 []byte, buffer2 []byte) ([]byte, error) {
	if len(buffer1) != len(buffer2) {
		return nil, errors.New("Buffers are not same size")
	}

	output := make([]byte, len(buffer1))
	for i := 0; i < len(buffer1); i++ {
		output[i] = buffer1[i] ^ buffer2[i]
	}
	return output, nil
}

func challenge2() {
	inputStr := "1c0111001f010100061a024b53535009181c"
	inputXorVal := "686974207468652062756c6c277320657965"
	desiredOutput := "746865206b696420646f6e277420706c6179"

	result, _ := xorTogether(hex_string_to_bytes(inputStr), hex_string_to_bytes(inputXorVal))
	passed := (hexBufferToString(result) == desiredOutput)

	if passed {
		fmt.Println("Challenge 2 Passed")
	} else {
		fmt.Println("Challenge 2 Failed")
	}
}

func challenge3() {
	// Approach: XOR with each possibility, score using occurances on alphabetical characters
	inputStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	dataBytes := hex_string_to_bytes(inputStr)

	letterCounts := make(map[byte]int)
	commonLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ '"

	for i := 0; i < 128; i++ {
		letterCounts[byte(i)] = 0
		for _, char := range commonLetters {
			letterCounts[byte(i)] += strings.Count(string(dataBytes), string(byte(char)^byte(i)))
		}
	}

	maxCount := 0
	maxOffset := byte(0)
	for key, value := range letterCounts {
		if value > maxCount {
			maxCount = value
			maxOffset = key
		}
	}
	for i := 0; i < len(dataBytes); i++ {
		dataBytes[i] ^= byte(maxOffset)
	}

	// Answer found using above code, but hardcoded here for regression detection purposes
	if string(dataBytes) == "Cooking MC's like a pound of bacon" {
		fmt.Println("Challenge 3 Passed")
	} else {
		fmt.Println("Challenge 3 Failed")
	}
}

func challenge4() {
	fileHandle, _ := os.Open("set1_q4_data.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	// Read lines nested byte arrays
	lines := make([][]byte, 0, 350)
	for fileScanner.Scan() {
		lines = append(lines, hex_string_to_bytes(fileScanner.Text()))
	}

	scores := make([]int, len(lines))

	// Construct a map containing the characters we care about for easy
	inAlphabet := make(map[byte]bool)
	symbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ '"
	for char := range symbols {
		char_byte := byte(char)

		inAlphabet[char_byte] = true
	}

	// Conduct a frequency analysis of each hex string to determine the outlier
	// Assume the other strings are all random

	fmt.Println(lines)
}

// General Usage

func hex_string_to_bytes(hex_s string) []byte {
	buffer, err := hex.DecodeString(hex_s)
	if err != nil {
		log.Fatal(err)
	}
	return buffer
}

func hexBufferToString(buffer []byte) string {
	return hex.EncodeToString(buffer)
}

func main() {
	challenge1()
	challenge2()
	challenge3()
	challenge4()
}
