package puzzlelib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"strings"
	"unicode"
	"unicode/utf8"

	"project/registry"
)

// StartJSON is the JSON start file handed to students.
type StartJSON struct {
	FirstCipher   string            `json:"first_cipher"`
	Params        map[string]string `json:"params,omitempty"`
	CiphertextB64 string            `json:"ciphertext_b64"`
}

// ClueJSON is the JSON produced after decrypting one layer.
type ClueJSON struct {
	Step          int               `json:"step,omitempty"`
	Total         int               `json:"total,omitempty"`
	NextCipher    string            `json:"next_cipher,omitempty"`
	NextParams    map[string]string `json:"next_params,omitempty"`
	CiphertextB64 string            `json:"ciphertext_b64,omitempty"`

	Done            bool   `json:"done,omitempty"`
	FinalMessageB64 string `json:"final_message_b64,omitempty"`
}

// Generate creates a nested clue chain that uses every registered cipher exactly once.
// finalMessage is the plaintext revealed after the last decrypt.
func Generate(finalMessage string, rng *rand.Rand) (start StartJSON, err error) {
	if rng == nil {
		rng = rand.New(rand.NewPCG(1, 2))
	}

	names := registry.CipherRegistry.Names()
	if len(names) == 0 {
		return StartJSON{}, fmt.Errorf("no ciphers registered")
	}

	// Shuffle for random order; students must implement all to proceed.
	rng.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	type step struct {
		Name   string
		Params map[string]string
	}
	steps := make([]step, 0, len(names))
	for _, name := range names {
		c, err := registry.CipherRegistry.Get(name)
		if err != nil {
			return StartJSON{}, err
		}
		steps = append(steps, step{
			Name:   c.Name(),
			Params: c.RandomParams(rng).Encode(),
		})
	}

	// Build nested JSON blobs from inner (final message) to outer (start blob).
	curPlain := []byte(finalMessage)
	for i := len(steps) - 1; i >= 0; i-- {
		hasNext := i < len(steps)-1
		var clue ClueJSON
		if hasNext {
			clue = ClueJSON{
				Step:          i + 1,
				Total:         len(steps),
				NextCipher:    steps[i+1].Name,
				NextParams:    steps[i+1].Params,
				CiphertextB64: base64.StdEncoding.EncodeToString(curPlain),
			}
		} else {
			clue = ClueJSON{
				Done:            true,
				FinalMessageB64: base64.StdEncoding.EncodeToString(curPlain),
			}
		}
		plainBytes, err := json.Marshal(clue)
		if err != nil {
			return StartJSON{}, err
		}

		c, err := registry.CipherRegistry.Get(steps[i].Name)
		if err != nil {
			return StartJSON{}, err
		}
		parsed, err := c.ParseParams(steps[i].Params)
		if err != nil {
			return StartJSON{}, err
		}
		enc, err := c.Encrypt(plainBytes, parsed)
		if err != nil {
			return StartJSON{}, err
		}
		curPlain = enc
	}

	start = StartJSON{
		FirstCipher:   steps[0].Name,
		Params:        steps[0].Params,
		CiphertextB64: base64.StdEncoding.EncodeToString(curPlain),
	}
	return start, nil
}

type SolveResult struct {
	StepsSolved int
	FinalText   string
}

// SolveFromStartText reads the plaintext start file and walks the chain until done.
func SolveFromStartText(startText string) (SolveResult, error) {
	return SolveFromStartTextTrace(startText, nil)
}

// SolveFromStartTextTrace behaves like SolveFromStartText but prints a step-by-step
// trace to traceW (typically os.Stderr). The final plaintext is still returned in
// SolveResult.
func SolveFromStartTextTrace(startText string, traceW io.Writer) (SolveResult, error) {
	trace := func(format string, args ...any) {
		if traceW == nil {
			return
		}
		_, _ = fmt.Fprintf(traceW, format, args...)
	}

	var start StartJSON
	if err := json.Unmarshal([]byte(startText), &start); err != nil {
		return SolveResult{}, fmt.Errorf("parse start JSON: %w", err)
	}
	if strings.TrimSpace(start.FirstCipher) == "" {
		return SolveResult{}, fmt.Errorf("start JSON missing first_cipher")
	}
	if strings.TrimSpace(start.CiphertextB64) == "" {
		return SolveResult{}, fmt.Errorf("start JSON missing ciphertext_b64")
	}
	if start.Params == nil {
		start.Params = map[string]string{}
	}

	curCipher := start.FirstCipher
	curParams := start.Params
	curB64 := start.CiphertextB64

	steps := 0
	for {
		trace("\n--- step %d ---\n", steps+1)
		trace("cipher: %s\n", curCipher)
		if len(curParams) == 0 {
			trace("params: (none)\n")
		} else {
			b, _ := json.MarshalIndent(curParams, "", "  ")
			trace("params: %s\n", string(b))
		}
		trace("input.ciphertext_b64: %s\n", curB64)

		ct, err := base64.StdEncoding.DecodeString(curB64)
		if err != nil {
			return SolveResult{}, fmt.Errorf("base64 decode: %w", err)
		}
		trace("input.ciphertext_len: %d bytes\n", len(ct))
		if s, ok := maybePrintableText(ct); ok {
			trace("input.ciphertext_text: %s\n", s)
		}

		c, err := registry.CipherRegistry.Get(curCipher)
		if err != nil {
			return SolveResult{}, err
		}
		parsed, err := c.ParseParams(curParams)
		if err != nil {
			return SolveResult{}, err
		}
		plain, err := c.Decrypt(ct, parsed)
		if err != nil {
			return SolveResult{}, err
		}
		trace("output.plaintext_len: %d bytes\n", len(plain))
		trace("output.plaintext_b64: %s\n", base64.StdEncoding.EncodeToString(plain))
		if s, ok := maybePrintableText(plain); ok {
			trace("output.plaintext_text:\n%s\n", s)
		}

		steps++

		var clue ClueJSON
		if err := json.Unmarshal(plain, &clue); err != nil {
			return SolveResult{}, fmt.Errorf("parse clue JSON at step %d: %w", steps, err)
		}
		pretty, _ := json.MarshalIndent(clue, "", "  ")
		trace("output.clue_json:\n%s\n", string(pretty))
		if clue.Done {
			finalBytes, err := base64.StdEncoding.DecodeString(clue.FinalMessageB64)
			if err != nil {
				return SolveResult{}, fmt.Errorf("decode final message: %w", err)
			}
			trace("\n--- done ---\n")
			trace("final_message_len: %d bytes\n", len(finalBytes))
			trace("final_message_b64: %s\n", base64.StdEncoding.EncodeToString(finalBytes))
			if s, ok := maybePrintableText(finalBytes); ok {
				trace("final_message_text:\n%s\n", s)
			}
			return SolveResult{StepsSolved: steps, FinalText: string(finalBytes)}, nil
		}
		if strings.TrimSpace(clue.NextCipher) == "" || strings.TrimSpace(clue.CiphertextB64) == "" {
			return SolveResult{}, fmt.Errorf("clue JSON missing next_cipher/ciphertext_b64 at step %d", steps)
		}
		if clue.NextParams == nil {
			clue.NextParams = map[string]string{}
		}
		curCipher = clue.NextCipher
		curParams = clue.NextParams
		curB64 = clue.CiphertextB64
	}
}

func maybePrintableText(b []byte) (string, bool) {
	if len(b) == 0 {
		return "", true
	}
	if !utf8.Valid(b) {
		return "", false
	}
	s := string(b)
	printable := 0
	non := 0
	for _, r := range s {
		switch {
		case r == '\n' || r == '\r' || r == '\t':
			printable++
		case unicode.IsPrint(r):
			printable++
		default:
			non++
		}
	}
	// Heuristic: show as text if overwhelmingly printable.
	return s, non == 0 || printable >= non*10
}

// Param generation is owned by each cipher (RandomParams), so tools do not need to
// understand parameter types, constraints, or names.
