package hashes

import (
	"testing"
)

func TestSHA256(t *testing.T) {
	got := SHA256("howsthisforatest")
	want := "d939b74aad696e1ece77ac6726bf9bad5a33d43cabf8c2df2ac7a44a0344dd66"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestMD5(t *testing.T) {
	got := MD5("another test to try")
	want := "26259b904807959e7559d67bb96a9eec"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
