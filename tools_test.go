package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	expectedLength := 10
	s := testTools.RandomString(expectedLength)
	if len(s) != expectedLength {
		t.Errorf("wrong length random string returned, expected length %d but got length %d", expectedLength, len(s))
	}
}
