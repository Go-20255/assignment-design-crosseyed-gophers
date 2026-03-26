package cipher

import "fmt"

/*
RailFenceCipher (technical specification)

- **Name**: `railfence`
- **Parameters**:
  - `--rails` (int, required): integer >= 2

- **Unit of operation**: runes (Unicode code points)
- **Encrypt**:
  - Write the input runes in a zig-zag down/up across `rails` rows:
    row sequence goes: 0,1,2,...,rails-1,rails-2,...,1,0,1,...
  - Ciphertext is the concatenation of row 0 then row 1 ... then row (rails-1)

- **Decrypt**:
  - Reconstruct the same zig-zag path indices for length N, split ciphertext into rows
    using per-row counts, then read off runes following the zig-zag path.
*/
type RailFenceCipher struct{}

func (RailFenceCipher) Name() string { return "railfence" }

func (RailFenceCipher) ParamSpecs() []ParamSpec {
	return []ParamSpec{{Name: "rails", Type: ParamInt, Required: true, Help: "Number of rails (>=2)"}}
}

type railFenceParams struct {
	rails int
}

func (railFenceParams) isParsedParams() {}

func (p railFenceParams) Encode() map[string]string {
	return map[string]string{"rails": encodeInt(p.rails)}
}

func (RailFenceCipher) ParseParams(raw map[string]string) (ParsedParams, error) {
	s, err := getRequired(raw, "rails")
	if err != nil {
		return nil, err
	}
	rails, err := parseIntParam(s)
	if err != nil {
		return nil, fmt.Errorf("rails: %w", err)
	}
	if rails < 2 {
		return nil, fmt.Errorf("rails must be >= 2 (got %d)", rails)
	}
	return railFenceParams{rails: rails}, nil
}

func (RailFenceCipher) RandomParams(rng Random) ParsedParams {
	rails := int(rng.Int32N(5)) + 2 // 2..6
	return railFenceParams{rails: rails}
}

func (RailFenceCipher) Encrypt(input []byte, params ParsedParams) ([]byte, error) {
	if _, ok := params.(railFenceParams); !ok {
		return nil, fmt.Errorf("railfence: wrong params type")
	}
	// STUDENT TODO: implement Rail Fence encryption (see spec above).
	_ = input
	return nil, ErrNotImplemented
}

func (RailFenceCipher) Decrypt(input []byte, params ParsedParams) ([]byte, error) {
	if _, ok := params.(railFenceParams); !ok {
		return nil, fmt.Errorf("railfence: wrong params type")
	}
	// STUDENT TODO: implement Rail Fence decryption.
	_ = input
	return nil, ErrNotImplemented
}
