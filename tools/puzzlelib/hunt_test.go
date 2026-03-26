package puzzlelib

import (
	"encoding/json"
	"math/rand/v2"
	"testing"
)

func TestGenerateThenSolveRoundTrip(t *testing.T) {
	rng := rand.New(rand.NewPCG(123, 456))
	final := "Skill issue."

	start, err := Generate(final, rng)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(start)
	if err != nil {
		t.Fatal(err)
	}

	res, err := SolveFromStartText(string(b))
	if err != nil {
		t.Fatal(err)
	}
	if res.FinalText != final {
		t.Fatalf("got %q want %q", res.FinalText, final)
	}
}
