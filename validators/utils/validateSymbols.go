package utils

func GetValidateSymbols() []string {
	// Добавляем символы, которые не разрешены
	return []string{"/", "*", "-", "+", "=", "_", "{", "}", "'", "\"", "[", "]", ";", ":", "#", "@", "!", "№", ",", ".", "$", "%", "^", "&", "?", "(", ")", "`"}
}

func GetAlphabetSymbols() map[rune]string {
	return map[rune]string{
		'a': "α", 'b': "β", 'c': "¢", 'd': "δ", 'e': "ε", 'f': "ƒ", 'g': "γ", 'h': "η", 'i': "ι", 'j': "ϑ",
		'k': "κ", 'l': "λ", 'm': "μ", 'n': "ν", 'o': "ο", 'p': "π", 'q': "θ", 'r': "ρ", 's': "σ", 't': "τ",
		'u': "υ", 'v': "ν", 'w': "ω", 'x': "χ", 'y': "ψ", 'z': "ζ",
	}
}
