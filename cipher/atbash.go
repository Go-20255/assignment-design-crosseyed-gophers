package cipher

/*
AtbashCipher (technical specification)

- **Name**: `atbash`
- **Parameters**: none
- **Alphabet**: ASCII letters only
- **Transform**:
  - Lowercase: `a` ↔ `z`, `b` ↔ `y`, ..., `m` ↔ `n`
  - Uppercase: `A` ↔ `Z`, `B` ↔ `Y`, ..., `M` ↔ `N`
  - Non-letters: unchanged

- **Symmetry**: encrypt == decrypt
*/
type AtbashCipher struct{}

func (AtbashCipher) Name() string { return "atbash" }

func (AtbashCipher) ParamSpecs() []ParamSpec { return nil }

type atbashParams struct{}

func (atbashParams) isParsedParams() {}

func (atbashParams) Encode() map[string]string { return map[string]string{} }

func (AtbashCipher) ParseParams(_ map[string]string) (ParsedParams, error) {
	return atbashParams{}, nil
}

func (AtbashCipher) RandomParams(_ Random) ParsedParams { return atbashParams{} }

func (AtbashCipher) Encrypt(input []byte, _ ParsedParams) ([]byte, error) {
	// STUDENT TODO: implement Atbash encryption (see spec above).
	_ = input
	return nil, ErrNotImplemented
}

func (AtbashCipher) Decrypt(input []byte, _ ParsedParams) ([]byte, error) {
	// STUDENT TODO: implement Atbash decryption.
	// NOTE: Atbash is symmetric, so this can call Encrypt.
	_ = input
	return nil, ErrNotImplemented
}
