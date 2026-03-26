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
type CaesarCipher struct {
}

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
	p, ok := params.(caesarParams)
	if !ok {
		return nil, fmt.Errorf("caesar: wrong params type")
	}
	return []byte(caesarEncrypt(string(input), p.shift)), nil
}

func (CaesarCipher) Decrypt(input []byte, params ParsedParams) ([]byte, error) {
	p, ok := params.(caesarParams)
	if !ok {
		return nil, fmt.Errorf("caesar: wrong params type")
	}
	return []byte(caesarDecrypt(string(input), p.shift)), nil
}

func caesarEncrypt(text string, shift int) string {
	result := ""
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			result += string((ch-'a'+rune(shift))%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			result += string((ch-'A'+rune(shift))%26 + 'A')
		} else {
			result += string(ch)
		}
	}
	return result
}

func caesarDecrypt(text string, shift int) string {
	result := ""
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			result += string((ch-'a'-rune(shift)+26)%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			result += string((ch-'A'-rune(shift)+26)%26 + 'A')
		} else {
			result += string(ch)
		}
	}
	return result
}
