package registry

import "project/cipher"

// CipherRegistry is the single registration point.
//
// STUDENT TODO:
//   - As you implement each cipher in `cipher/`, add it here.
//   - Only ciphers registered here will show up in the CLI and be usable
var CipherRegistry = cipher.NewRegistry(
	cipher.CaesarCipher{},
	// oops, it looks like one was left here...
	// use this as an example to register other ciphers
)
