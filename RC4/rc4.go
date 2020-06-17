package rc4

const sboxSize int = 256

// GenerateKeyStream generates a keystream of streamLen bytes from a key of 40, 56, 64, 80, 128, 192, or 256 bits bits length
func GenerateKeyStream(key []byte, streamLen int) []byte {
	keybuf := make([]byte, len(key), len(key))
	for i := range key {
		keybuf[i] = byte(key[i])
	}
	sbox := make([]byte, sboxSize, sboxSize)
	for i := range sbox {
		sbox[i] = byte(i)
	}

	K := make([]byte, sboxSize, sboxSize)
	for i := range K {
		K[i] = keybuf[i%len(key)]
	}

	var j uint8 = 0
	for i := range sbox {
		j = j + uint8(sbox[i]) + uint8(K[i]) // should overflow and keep mod256

		// swap si and sj
		temp := sbox[i]
		sbox[i] = sbox[j]
		sbox[j] = temp
	}

	keyStream := make([]byte, streamLen, streamLen)
	j = 0
	var i uint8 = 0
	var t uint8 = 0

	for k := range keyStream {
		i = i + 1
		j = j + uint8(sbox[i])

		temp := sbox[i]
		sbox[i] = sbox[j]
		sbox[j] = temp

		t = uint8(sbox[i]) + uint8(sbox[j])

		keyStream[k] = sbox[t]
	}

	return keyStream
}

// TODO encryption (trivial since keystream is already done)
// Add configurable offset as in RC4 spec
