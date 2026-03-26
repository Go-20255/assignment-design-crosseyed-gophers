package cipher

import "errors"

// ErrNotImplemented is returned by student stubs until a cipher is completed.
var ErrNotImplemented = errors.New("cipher: not implemented")

// Cipher is the common interface all ciphers implement.
// Each cipher is the single source of truth for:
// - its name
// - its parameter schema
// - encrypt/decrypt behavior
type Cipher interface {
	Name() string
	ParamSpecs() []ParamSpec
	ParseParams(raw map[string]string) (ParsedParams, error)
	RandomParams(rng Random) ParsedParams
	Encrypt(input []byte, params ParsedParams) ([]byte, error)
	Decrypt(input []byte, params ParsedParams) ([]byte, error)
}

// ParamType describes the type a CLI parameter expects.
type ParamType string

const (
	ParamInt    ParamType = "int"
	ParamString ParamType = "string"
	ParamByte   ParamType = "byte" // integer 0-255
)

// ParamSpec is owned by the cipher and describes one CLI parameter.
type ParamSpec struct {
	Name     string
	Type     ParamType
	Required bool
	Help     string
}

// (No separate Definition type: the cipher itself owns schema + parsing.)

// Random is the minimal RNG surface ciphers need for RandomParams.
// It is satisfied by math/rand/v2.Rand.
type Random interface {
	Int32N(n int32) int32
}

// ParsedParams is an opaque, typed parameter value produced by a cipher's ParseParams
// (or RandomParams). Invalid states should be unrepresentable: callers should never
// construct ParsedParams themselves.
type ParsedParams interface {
	isParsedParams()
	Encode() map[string]string
}
