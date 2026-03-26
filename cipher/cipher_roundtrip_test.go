package cipher

import (
	"bytes"
	"testing"
)

func TestCipherRoundTrip(t *testing.T) {
	input := []byte("hello world this is secret")

	cases := []struct {
		name   string
		params map[string]string
	}{
		{"caesar", map[string]string{"shift": "3"}},
		{"atbash", map[string]string{}},
		{"vigenere", map[string]string{"key": "KEY"}},
		{"railfence", map[string]string{"rails": "3"}},
		{"xor", map[string]string{"key": "5"}},
		{"substitution", map[string]string{"key": "qwertyuiopasdfghjklzxcvbnm"}},
	}

	for _, tc := range cases {
		c, err := testRegistry().Get(tc.name)
		if err != nil {
			t.Fatalf("%s: get cipher: %v", tc.name, err)
		}

		parsed, err := c.ParseParams(tc.params)
		if err != nil {
			t.Fatalf("%s: parse params: %v", tc.name, err)
		}

		enc, err := c.Encrypt(input, parsed)
		if err != nil {
			t.Fatalf("%s: encrypt: %v", tc.name, err)
		}
		dec, err := c.Decrypt(enc, parsed)
		if err != nil {
			t.Fatalf("%s: decrypt: %v", tc.name, err)
		}

		if !bytes.Equal(dec, input) {
			t.Fatalf("%s: round-trip mismatch\nwant: %q\ngot:  %q", tc.name, input, dec)
		}
	}
}

func TestCipherKnownVectors(t *testing.T) {
	r := testRegistry()

	t.Run("atbash", func(t *testing.T) {
		c, _ := r.Get("atbash")
		p, err := c.ParseParams(map[string]string{})
		if err != nil {
			t.Fatal(err)
		}
		out, err := c.Encrypt([]byte("abC-xyz"), p)
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != "zyX-cba" {
			t.Fatalf("got %q", string(out))
		}
	})

	t.Run("caesar", func(t *testing.T) {
		c, _ := r.Get("caesar")
		p, err := c.ParseParams(map[string]string{"shift": "3"})
		if err != nil {
			t.Fatal(err)
		}
		out, err := c.Encrypt([]byte("Abc XyZ!"), p)
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != "Def AbC!" {
			t.Fatalf("got %q", string(out))
		}
	})

	t.Run("vigenere", func(t *testing.T) {
		c, _ := r.Get("vigenere")
		p, err := c.ParseParams(map[string]string{"key": "LEMON"})
		if err != nil {
			t.Fatal(err)
		}
		out, err := c.Encrypt([]byte("ATTACKATDAWN"), p)
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != "LXFOPVEFRNHR" {
			t.Fatalf("got %q", string(out))
		}
		plain, err := c.Decrypt(out, p)
		if err != nil {
			t.Fatal(err)
		}
		if string(plain) != "ATTACKATDAWN" {
			t.Fatalf("got %q", string(plain))
		}
	})

	t.Run("railfence", func(t *testing.T) {
		c, _ := r.Get("railfence")
		p, err := c.ParseParams(map[string]string{"rails": "3"})
		if err != nil {
			t.Fatal(err)
		}
		out, err := c.Encrypt([]byte("WEAREDISCOVEREDFLEEATONCE"), p)
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != "WECRLTEERDSOEEFEAOCAIVDEN" {
			t.Fatalf("got %q", string(out))
		}
	})

	t.Run("substitution", func(t *testing.T) {
		c, _ := r.Get("substitution")
		p, err := c.ParseParams(map[string]string{"key": "qwertyuiopasdfghjklzxcvbnm"})
		if err != nil {
			t.Fatal(err)
		}
		out, err := c.Encrypt([]byte("abc-xyz"), p)
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != "qwe-bnm" {
			t.Fatalf("got %q", string(out))
		}
	})

	t.Run("xor", func(t *testing.T) {
		c, _ := r.Get("xor")
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
	})
}

func testRegistry() *Registry {
	return NewRegistry(
		AtbashCipher{},
		CaesarCipher{},
		VigenereCipher{},
		RailFenceCipher{},
		XORCipher{},
		SubstitutionCipher{},
	)
}
