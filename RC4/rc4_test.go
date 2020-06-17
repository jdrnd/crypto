package rc4

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"testing"
)

func readTestVectors() ([]string, []string) {
	file, err := os.Open("vectors.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanBuf := make([]byte, 1e6)
	scanner.Buffer(scanBuf, len(scanBuf))

	keys := make([]string, 0, 15)
	vectors := make([]string, 0, 15)

	keyv := true
	for scanner.Scan() {
		line := scanner.Text()

		if keyv == true {
			keys = append(keys, line)
			keyv = false
		} else {
			vectors = append(vectors, line)
			keyv = true
		}
	}

	return keys, vectors
}

func TestRC4Vectors(t *testing.T) {
	fmt.Println("Testing against known test vectors")
	keys, vectors := readTestVectors()

	for i := range keys {
		keystr := []byte(keys[i])
		key := make([]byte, hex.DecodedLen(len(keystr)))
		vectorstr := []byte(vectors[i])
		vector := make([]byte, hex.DecodedLen(len(vectorstr)))

		hex.Decode(key, keystr)
		hex.Decode(vector, vectorstr)

		passed := bytes.Equal(GenerateKeyStream(key, 32), vector[:32])

		if !passed {
			t.Errorf("Output did not match test vector")
		}
	}
	k1 := []byte("0102030405")
	dst := make([]byte, hex.DecodedLen(len(k1)))

	hex.Decode(dst, k1)
	fmt.Println(hex.Dump(GenerateKeyStream(dst, 256)))
}

// TODO add a benchmark for speed
