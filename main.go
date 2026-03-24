package main

import (
	"fmt"
	"project/decrypt"
	"project/encrypt"
)

func main() {
	text := encrypt.ReadFile("test/test_1.txt")
	fmt.Println("Original Text:", text)

	// 1. Caesar
	caesar := encrypt.CaesarEncrypt(text, 3)
	fmt.Println("\nCaesar Encrypted:", caesar)
	fmt.Println("Caesar Decrypted:", decrypt.CaesarDecrypt(caesar, 3))

	// 2. Atbash
	atbash := encrypt.AtbashEncrypt(text)
	fmt.Println("\nAtbash Encrypted:", atbash)
	fmt.Println("Atbash Decrypted:", decrypt.AtbashDecrypt(atbash))

	// 3. Vigenere
	vigenere := encrypt.VigenereEncrypt(text, "KEY")
	fmt.Println("\nVigenere Encrypted:", vigenere)
	fmt.Println("Vigenere Decrypted:", decrypt.VigenereDecrypt(vigenere, "KEY"))

	// 4. Rail Fence
	rail := encrypt.RailFenceEncrypt(text, 3)
	fmt.Println("\nRail Fence Encrypted:", rail)
	fmt.Println("Rail Fence Decrypted:", decrypt.RailFenceDecrypt(rail, 3))

	// 5. XOR
	xor := encrypt.XOREncrypt(text, 5)
	fmt.Println("\nXOR Encrypted:", xor)
	fmt.Println("XOR Decrypted:", decrypt.XORDecrypt(xor, 5))

	// 6. Substitution
	sub := encrypt.SubstitutionEncrypt(text)
	fmt.Println("\nSubstitution Encrypted:", sub)
	fmt.Println("Substitution Decrypted:", decrypt.SubstitutionDecrypt(sub))
}
