package cipher

import "testing"

func TestRailFenceCipher_Vector(t *testing.T) {
	c := RailFenceCipher{}
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
}
