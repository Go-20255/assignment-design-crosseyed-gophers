package cipher

import "fmt"

/*
SubstitutionCipher (technical specification)

- **Name**: `substitution`
- **Parameters**:
  - `--key` (string, required): a 26-character permutation of `a`..`z`

- **Alphabet**: lowercase ASCII letters `a`..`z` only
- **Encrypt**:
  - For each rune `ch` in the input:
  - if `ch` is in `a`..`z`, replace with `key[ch-'a']`
  - otherwise, leave unchanged

- **Decrypt**:
  - same rule using the inverse mapping induced by `key`

Example key (classic “qwerty” mapping, used in our tests):

	`qwertyuiopasdfghjklzxcvbnm`
*/
type SubstitutionCipher struct{}

func (SubstitutionCipher) Name() string { return "substitution" }

func (SubstitutionCipher) ParamSpecs() []ParamSpec {
	return []ParamSpec{{
		Name:     "key",
		Type:     ParamString,
		Required: true,
		Help:     "26-letter permutation of a-z (e.g. qwertyuiopasdfghjklzxcvbnm)",
	}}
}

type substitutionParams struct {
	key string
	enc [26]rune
	dec [26]rune
}

func (substitutionParams) isParsedParams() {}

func (p substitutionParams) Encode() map[string]string { return map[string]string{"key": p.key} }

func (SubstitutionCipher) ParseParams(raw map[string]string) (ParsedParams, error) {
	key, err := getRequired(raw, "key")
	if err != nil {
		return nil, err
	}
	p, err := parseSubstitutionKey(key)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (SubstitutionCipher) RandomParams(rng Random) ParsedParams {
	// Fisher–Yates shuffle of 'a'..'z'
	b := make([]byte, 26)
	for i := 0; i < 26; i++ {
		b[i] = byte('a' + i)
	}
	for i := 25; i > 0; i-- {
		j := int(rng.Int32N(int32(i + 1)))
		b[i], b[j] = b[j], b[i]
	}
	p, _ := parseSubstitutionKey(string(b))
	return p
}

func (SubstitutionCipher) Encrypt(input []byte, params ParsedParams) ([]byte, error) {
	p, ok := params.(substitutionParams)
	if !ok {
		return nil, fmt.Errorf("substitution: wrong params type")
	}
	return []byte(substitutionEncrypt(string(input), p.enc)), nil
}

func (SubstitutionCipher) Decrypt(input []byte, params ParsedParams) ([]byte, error) {
	p, ok := params.(substitutionParams)
	if !ok {
		return nil, fmt.Errorf("substitution: wrong params type")
	}
	return []byte(substitutionDecrypt(string(input), p.dec)), nil
}

func parseSubstitutionKey(key string) (substitutionParams, error) {
	if len(key) != 26 {
		return substitutionParams{}, fmt.Errorf("key must be 26 characters (got %d)", len(key))
	}

	var enc [26]rune
	var dec [26]rune
	var seen [26]bool

	for i := 0; i < 26; i++ {
		ch := rune(key[i])
		if ch < 'a' || ch > 'z' {
			return substitutionParams{}, fmt.Errorf("key must contain only lowercase a-z (bad char %q at index %d)", ch, i)
		}
		idx := int(ch - 'a')
		if seen[idx] {
			return substitutionParams{}, fmt.Errorf("key must be a permutation of a-z (duplicate %q)", ch)
		}
		seen[idx] = true
		enc[i] = ch
		dec[idx] = rune('a' + i)
	}

	return substitutionParams{key: key, enc: enc, dec: dec}, nil
}

func substitutionEncrypt(text string, enc [26]rune) string {
	result := ""
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			result += string(enc[ch-'a'])
		} else {
			result += string(ch)
		}
	}
	return result
}

func substitutionDecrypt(text string, dec [26]rune) string {
	result := ""
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			result += string(dec[ch-'a'])
		} else {
			result += string(ch)
		}
	}
	return result
}
