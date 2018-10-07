package cryptopals

import (
	"strings"
	"testing"
)

func TestSet1Challenge7(t *testing.T) {
	input_file := "../inputs/7.txt"
	out := set1Challenge7(input_file)

	if !strings.Contains(out, "There's no trippin' on mine, I'm just gettin' down") {
		t.Log(out)
		t.Error("Did not decrypt ECB mode correctly")
	}
}

func TestSet1Challenge8(t *testing.T) {}
