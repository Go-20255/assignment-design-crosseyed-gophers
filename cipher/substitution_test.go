package cipher

import "testing"

func TestSubstitutionCipher_Vector(t *testing.T) {
	c := SubstitutionCipher{}
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
}
