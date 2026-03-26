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
type RailFenceCipher struct {
}

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
	p, ok := params.(railFenceParams)
	if !ok {
		return nil, fmt.Errorf("railfence: wrong params type")
	}
	return []byte(railFenceEncrypt(string(input), p.rails)), nil
}

func (RailFenceCipher) Decrypt(input []byte, params ParsedParams) ([]byte, error) {
	p, ok := params.(railFenceParams)
	if !ok {
		return nil, fmt.Errorf("railfence: wrong params type")
	}
	return []byte(railFenceDecrypt(string(input), p.rails)), nil
}

func railFenceEncrypt(text string, rails int) string {
	if rails <= 1 {
		return text
	}

	rows := make([][]rune, rails)
	dir := 1
	row := 0
	for _, ch := range text {
		rows[row] = append(rows[row], ch)
		row += dir
		if row == 0 || row == rails-1 {
			dir *= -1
		}
	}

	result := ""
	for _, r := range rows {
		result += string(r)
	}
	return result
}

func railFenceDecrypt(text string, rails int) string {
	if rails <= 1 {
		return text
	}

	runes := []rune(text)
	n := len(runes)
	rowsIdx := make([]int, n)

	dir := 1
	row := 0
	for i := 0; i < n; i++ {
		rowsIdx[i] = row
		row += dir
		if row == 0 || row == rails-1 {
			dir *= -1
		}
	}

	count := make([]int, rails)
	for _, r := range rowsIdx {
		count[r]++
	}

	resultRows := make([][]rune, rails)
	idx := 0
	for r := 0; r < rails; r++ {
		resultRows[r] = append([]rune(nil), runes[idx:idx+count[r]]...)
		idx += count[r]
	}

	pos := make([]int, rails)
	dir = 1
	row = 0
	out := make([]rune, n)
	for i := 0; i < n; i++ {
		out[i] = resultRows[row][pos[row]]
		pos[row]++

		row += dir
		if row == 0 || row == rails-1 {
			dir *= -1
		}
	}

	return string(out)
}
