// Package cipher implements pluggable classical ciphers behind a common interface.
//
// Reference implementation: ciphers are wired into the central registry in the root.
//
// # Caesar
//
// Shift letters in A–Z and a–z by Shift positions, wrapping within each alphabet.
// Other bytes (spaces, punctuation) are unchanged.
//
// # Atbash
//
// Map A↔Z and a↔z within each alphabet; symmetric (encrypt == decrypt).
//
// # Vigenère
//
// Key repeats; only advances on letters. Non-letters pass through unchanged.
//
// # Rail Fence
//
// Zig-zag write across `Rails` rows, then read rows top-to-bottom.
// Decrypt reconstructs the zig-zag read order (length-preserving).
//
// # XOR
//
// XOR every byte with a single-byte key (encrypt == decrypt).
//
// # Substitution
//
// Use SubstitutionMap for lower-case a–z only; build the inverse map for decrypt.

package cipher
