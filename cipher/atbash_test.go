package cipher

import "testing"

func TestAtbashCipher_Vector(t *testing.T) {
	c := AtbashCipher{}
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
}
