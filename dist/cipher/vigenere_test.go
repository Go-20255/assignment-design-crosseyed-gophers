package cipher

import "testing"

func TestVigenereCipher_Vector(t *testing.T) {
	c := VigenereCipher{}
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
}
