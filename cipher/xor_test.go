package cipher

import (
	"bytes"
	"testing"
)

func TestXORCipher_Vector(t *testing.T) {
	c := XORCipher{}
	p, err := c.ParseParams(map[string]string{"key": "255"})
	if err != nil {
		t.Fatal(err)
	}
	in := []byte{0x00, 0x01, 0xAB, 0xFF}
	out, err := c.Encrypt(in, p)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{0xFF, 0xFE, 0x54, 0x00}
	if !bytes.Equal(out, want) {
		t.Fatalf("got %v want %v", out, want)
	}
}
