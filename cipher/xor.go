package cipher

import "fmt"

/*
XORCipher (technical specification)

- **Name**: `xor`
- **Parameters**:
  - `--key` (byte, required): integer in [0,255]

- **Unit of operation**: bytes
- **Encrypt**:
  - For each byte `b` in input: output byte is `b ^ key`

- **Decrypt**: identical to encrypt (XOR is its own inverse)
*/
type XORCipher struct {
}

func (XORCipher) Name() string { return "xor" }

func (XORCipher) ParamSpecs() []ParamSpec {
	return []ParamSpec{{Name: "key", Type: ParamByte, Required: true, Help: "XOR key byte (0-255)"}}
}

type xorParams struct {
	key byte
}

func (xorParams) isParsedParams() {}

func (p xorParams) Encode() map[string]string { return map[string]string{"key": encodeByte(p.key)} }

func (XORCipher) ParseParams(raw map[string]string) (ParsedParams, error) {
	s, err := getRequired(raw, "key")
	if err != nil {
		return nil, err
	}
	k, err := parseByteParam(s)
	if err != nil {
		return nil, fmt.Errorf("key: %w", err)
	}
	return xorParams{key: k}, nil
}

func (XORCipher) RandomParams(rng Random) ParsedParams {
	k := byte(rng.Int32N(255) + 1) // 1..255
	return xorParams{key: k}
}

func (XORCipher) Encrypt(input []byte, params ParsedParams) ([]byte, error) {
	if input == nil {
		return nil, fmt.Errorf("xor: input must not be nil")
	}
	p, ok := params.(xorParams)
	if !ok {
		return nil, fmt.Errorf("xor: wrong params type")
	}
	return xorBytes(input, p.key), nil
}

func (XORCipher) Decrypt(input []byte, params ParsedParams) ([]byte, error) {
	if input == nil {
		return nil, fmt.Errorf("xor: input must not be nil")
	}
	p, ok := params.(xorParams)
	if !ok {
		return nil, fmt.Errorf("xor: wrong params type")
	}
	return xorBytes(input, p.key), nil
}

func xorBytes(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key
	}
	return result
}
