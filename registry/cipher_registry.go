package registry

import "project/cipher"

// CipherRegistry is the single wiring point for all ciphers in this project.
// CLI and tools depend on this registry; the registry depends on the cipher implementations.
var CipherRegistry = cipher.NewRegistry(
	cipher.AtbashCipher{},
	cipher.CaesarCipher{},
	cipher.VigenereCipher{},
	cipher.RailFenceCipher{},
	cipher.XORCipher{},
	cipher.SubstitutionCipher{},
)
