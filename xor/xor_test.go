package crypto

import (
	"io/ioutil"
	"testing"
)

const input_file = "../inputs/plaintext.txt"

func TestBasicEnc(t *testing.T) {
	null_str := string([]byte{0x00, 0x00, 0x00, 0x00, 0x00})
	data := encrypt_string(null_str, "stuff")
	if string(data) != "stuff" {
		t.Errorf("Null string with key %s encrypted to %s", "stuff", string(data))
	}
}

func TestBasicDec(t *testing.T) {
	encrypted := "stuffstuff"
	data := decrypt_string(encrypted, "stuff")
	if string(data) != string([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) {
		t.Errorf("Encrypted string %s with key \"stuff\" decrypted to %s", string(encrypted), string(data))
	}
}

func TestKeyLength(t *testing.T) {
	data := encrypt_file(input_file, "test")
	key_len := get_key_length(data)
	if key_len != 4 {
		t.Log(key_len)
		t.Error("incorrect key length returned")
	}

	data = encrypt_file(input_file, "longerpassword")
	key_len = get_key_length(data)
	if key_len != 14 {
		t.Log(key_len)
		t.Error("incorrect key length returned")
	}

	data = encrypt_file(input_file, "longerpasswordwithmorestuff")
	key_len = get_key_length(data)
	if key_len != 27 {
		t.Log(key_len)
		t.Error("incorrect key length returned")
	}
}

func TestBreakingCypher(t *testing.T) {
	data := encrypt_file(input_file, "testingpass")
	key := string(get_multi_character_xor_key(data, get_key_length(data)))

	if key != "testingpass" {
		t.Log(key)
		t.Error("Incorect key found")
	}

	plaintext := break_cypher(string(data))
	example_plaintext, _ := ioutil.ReadFile(input_file)

	if plaintext != string(example_plaintext) {
		t.Log(string(plaintext))
		t.Error("Plaintext incorrectly decrypted with calculated key")
	}
}
