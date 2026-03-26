package cipher

import "testing"

func TestCaesarCipher_Vector(t *testing.T) {
	c := CaesarCipher{}
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
}
