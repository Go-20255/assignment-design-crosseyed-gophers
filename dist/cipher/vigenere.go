package cipher

import (
	"fmt"
	"strings"
)

/*
VigenereCipher (technical specification)

- **Name**: `vigenere`
- **Parameters**:
  - `--key` (string, required): non-empty string

- **Alphabet**: ASCII letters only
- **Key schedule**:
  - The key index increments **only when processing a letter** in the input
  - For each processed letter, take `k = key[keyIndex % len(key)]`
  - Convert `k` to a shift in [0,25]:
  - if `k` in `a`..`z`, shift = k-'a'
  - else (expected `A`..`Z`), shift = k-'A'

- **Encrypt**:
  - For lowercase letters: shift forward by `shift` modulo 26
  - For uppercase letters: shift forward by `shift` modulo 26
  - Non-letters: unchanged (and do not consume key)

- **Decrypt**: inverse shift (shift backward by `shift` modulo 26)

Note: This spec intentionally matches our implementation (we only validate non-empty;
we do not require the key to be letters).
*/
type VigenereCipher struct{}

func (VigenereCipher) Name() string { return "vigenere" }

func (VigenereCipher) ParamSpecs() []ParamSpec {
	return []ParamSpec{{Name: "key", Type: ParamString, Required: true, Help: "Key (A-Z letters recommended)"}}
}

type vigenereParams struct {
	key string
}

func (vigenereParams) isParsedParams() {}

func (p vigenereParams) Encode() map[string]string { return map[string]string{"key": p.key} }

func (VigenereCipher) ParseParams(raw map[string]string) (ParsedParams, error) {
	key, err := getRequired(raw, "key")
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(key) == "" {
		return nil, fmt.Errorf("key must be non-empty")
	}
	return vigenereParams{key: key}, nil
}

func (VigenereCipher) RandomParams(rng Random) ParsedParams {
	n := int(rng.Int32N(5)) + 2 // 2..6
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte('A' + rng.Int32N(26))
	}
	return vigenereParams{key: string(b)}
}

func (VigenereCipher) Encrypt(input []byte, params ParsedParams) ([]byte, error) {
	if _, ok := params.(vigenereParams); !ok {
		return nil, fmt.Errorf("vigenere: wrong params type")
	}
	// STUDENT TODO: implement Vigenère encryption (see spec above).
	_ = input
	return nil, ErrNotImplemented
}

func (VigenereCipher) Decrypt(input []byte, params ParsedParams) ([]byte, error) {
	if _, ok := params.(vigenereParams); !ok {
		return nil, fmt.Errorf("vigenere: wrong params type")
	}
	// STUDENT TODO: implement Vigenère decryption.
	_ = input
	return nil, ErrNotImplemented
}
