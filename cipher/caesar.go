package cipher

import "fmt"

/*
CaesarCipher (technical specification)

- **Name**: `caesar`
- **Parameters**:
  - `--shift` (int, required): an integer in [1, 25]

- **Alphabet**: ASCII letters only
- **Encrypt**:
  - For each lowercase letter `a`..`z`, shift forward by `shift` modulo 26
  - For each uppercase letter `A`..`Z`, shift forward by `shift` modulo 26
  - Non-letters: unchanged

- **Decrypt**: inverse shift (shift backward by `shift` modulo 26)
*/
type CaesarCipher struct{}

func (CaesarCipher) Name() string { return "caesar" }

func (CaesarCipher) ParamSpecs() []ParamSpec {
	return []ParamSpec{{Name: "shift", Type: ParamInt, Required: true, Help: "Shift amount (1-25)"}}
}

type caesarParams struct {
	shift int
}

func (caesarParams) isParsedParams() {}

func (p caesarParams) Encode() map[string]string {
	return map[string]string{"shift": encodeInt(p.shift)}
}

func (CaesarCipher) ParseParams(raw map[string]string) (ParsedParams, error) {
	s, err := getRequired(raw, "shift")
	if err != nil {
		return nil, err
	}
	shift, err := parseIntParam(s)
	if err != nil {
		return nil, fmt.Errorf("shift: %w", err)
	}
	if shift < 1 || shift > 25 {
		return nil, fmt.Errorf("shift must be in [1,25] (got %d)", shift)
	}
	return caesarParams{shift: shift}, nil
}

func (CaesarCipher) RandomParams(rng Random) ParsedParams {
	shift := int(rng.Int32N(25)) + 1
	return caesarParams{shift: shift}
}

func (CaesarCipher) Encrypt(input []byte, params ParsedParams) ([]byte, error) {
	_, ok := params.(caesarParams)
	if !ok {
		return nil, fmt.Errorf("caesar: wrong params type")
	}
	// STUDENT TODO: implement Caesar encryption.
	// Hints:
	// - Only shift ASCII letters A-Z and a-z
	// - Wrap within the alphabet (mod 26)
	// - Leave all other bytes unchanged
	_ = input
	return nil, ErrNotImplemented
}

func (CaesarCipher) Decrypt(input []byte, params ParsedParams) ([]byte, error) {
	_, ok := params.(caesarParams)
	if !ok {
		return nil, fmt.Errorf("caesar: wrong params type")
	}
	// STUDENT TODO: implement Caesar decryption (inverse of Encrypt).
	_ = input
	return nil, ErrNotImplemented
}
