package encrypt

import "os"

// --------------------- 1. Caesar Cipher ---------------------
func CaesarEncrypt(text string, shift int) string {
	result := ""
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			result += string((ch-'a'+rune(shift))%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			result += string((ch-'A'+rune(shift))%26 + 'A')
		} else {
			result += string(ch)
		}
	}
	return result
}

// --------------------- 2. Atbash Cipher ---------------------
func AtbashEncrypt(text string) string {
	result := ""
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			result += string('z' - (ch - 'a'))
		} else if ch >= 'A' && ch <= 'Z' {
			result += string('Z' - (ch - 'A'))
		} else {
			result += string(ch)
		}
	}
	return result
}

// --------------------- 3. Vigenere Cipher ---------------------
func VigenereEncrypt(text, key string) string {
	result := ""
	keyLen := len(key)
	keyIndex := 0

	for _, ch := range text {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			// convert key letter to 0-25 number
			k := rune(key[keyIndex%keyLen])
			if k >= 'a' && k <= 'z' {
				k -= 'a'
			} else {
				k -= 'A'
			}

			var base rune
			if ch >= 'a' && ch <= 'z' {
				base = 'a'
			} else {
				base = 'A'
			}
			result += string((ch-base+k)%26 + base)
			keyIndex++ // only advance key for letters
		} else {
			result += string(ch)
		}
	}
	return result
}

// --------------------- 4. Rail Fence Cipher ---------------------
func RailFenceEncrypt(text string, rails int) string {
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

// --------------------- 5. XOR Cipher ---------------------
func XOREncrypt(text string, key byte) []byte {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return result
}

// --------------------- Helper: Read file ---------------------
func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
