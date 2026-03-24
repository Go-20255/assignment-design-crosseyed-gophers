package encrypt

// Exported Substitution Map
var SubstitutionMap = map[rune]rune{
	'a': 'q', 'b': 'w', 'c': 'e', 'd': 'r', 'e': 't', 'f': 'y', 'g': 'u', 'h': 'i',
	'i': 'o', 'j': 'p', 'k': 'a', 'l': 's', 'm': 'd', 'n': 'f', 'o': 'g', 'p': 'h',
	'q': 'j', 'r': 'k', 's': 'l', 't': 'z', 'u': 'x', 'v': 'c', 'w': 'v', 'x': 'b',
	'y': 'n', 'z': 'm',
}

func SubstitutionEncrypt(text string) string {
	result := ""
	for _, ch := range text {
		if val, ok := SubstitutionMap[ch]; ok {
			result += string(val)
		} else {
			result += string(ch)
		}
	}
	return result
}
