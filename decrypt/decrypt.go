package decrypt

import "project/encrypt"

// --------------------- 1. Caesar Cipher ---------------------
func CaesarDecrypt(text string, shift int) string {
	result := ""
	for _, ch := range text {
		if ch >= 'a' && ch <= 'z' {
			result += string((ch-'a'-rune(shift)+26)%26 + 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			result += string((ch-'A'-rune(shift)+26)%26 + 'A')
		} else {
			result += string(ch)
		}
	}
	return result
}

// --------------------- 2. Atbash Cipher ---------------------
func AtbashDecrypt(text string) string {
	// symmetric cipher
	return encrypt.AtbashEncrypt(text)
}

// --------------------- 3. Vigenere Cipher ---------------------
func VigenereDecrypt(text, key string) string {
	result := ""
	keyLen := len(key)
	keyIndex := 0

	for _, ch := range text {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
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
			result += string((ch-base-k+26)%26 + base)
			keyIndex++
		} else {
			result += string(ch)
		}
	}
	return result
}

// --------------------- 4. Rail Fence Cipher ---------------------
func RailFenceDecrypt(text string, rails int) string {
	if rails <= 1 {
		return text
	}
	n := len(text)
	rows := make([]int, n)
	dir := 1
	row := 0
	for i := 0; i < n; i++ {
		rows[i] = row
		row += dir
		if row == 0 || row == rails-1 {
			dir *= -1
		}
	}
	count := make([]int, rails)
	for _, r := range rows {
		count[r]++
	}
	resultRows := make([][]rune, rails)
	idx := 0
	for r := 0; r < rails; r++ {
		resultRows[r] = []rune(text[idx : idx+count[r]])
		idx += count[r]
	}
	pos := make([]int, rails)
	row = 0
	dir = 1
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

// --------------------- 5. XOR Cipher ---------------------
func XORDecrypt(data []byte, key byte) []byte {
	// XOR decryption = encryption
	return XOREncrypt(data, key)
}

func XOREncrypt(data []byte, key byte) []byte {
	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ key
	}
	return result
}

// --------------------- 6. Substitution Cipher ---------------------
var reverseSubMap = func() map[rune]rune {
	m := make(map[rune]rune)
	for k, v := range encrypt.SubstitutionMap {
		m[v] = k
	}
	return m
}()

func SubstitutionDecrypt(text string) string {
	result := ""
	for _, ch := range text {
		if val, ok := reverseSubMap[ch]; ok {
			result += string(val)
		} else {
			result += string(ch)
		}
	}
	return result
}
