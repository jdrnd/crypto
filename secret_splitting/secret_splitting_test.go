package secret_splitting

import (
	"testing"
)

func TestSplitSecret(t *testing.T) {
	n := 2
	secret := "testing1234"
	result := split_secret(secret, n)
	t.Logf("Split '%s' into %s", secret, result)
	if len(result) != 2 {
		t.Errorf("Incorrect number of parts returned, %d returned, expected %d", len(result), n)
	}
	recombined := combine_secret(result)
	t.Logf("Combined %s into '%s'", result, secret)

	if recombined != secret {
		t.Errorf("Recombined %s into %s instead of %s", result, recombined, secret)
	}

}
