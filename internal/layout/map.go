package layout

// Bidirectional ЙЦУКЕН (Windows) ↔ QWERTY mapping.
var enRu = map[rune]rune{
	'q': 'й', 'w': 'ц', 'e': 'у', 'r': 'к', 't': 'е', 'y': 'н', 'u': 'г', 'i': 'ш', 'o': 'щ', 'p': 'з',
	'[': 'х', ']': 'ъ', 'a': 'ф', 's': 'ы', 'd': 'в', 'f': 'а', 'g': 'п', 'h': 'р', 'j': 'о', 'k': 'л', 'l': 'д',
	';': 'ж', '\'': 'э', 'z': 'я', 'x': 'ч', 'c': 'с', 'v': 'м', 'b': 'и', 'n': 'т', 'm': 'ь',
	',': 'б', '.': 'ю', '/': '.', '`': 'ё',
	'Q': 'Й', 'W': 'Ц', 'E': 'У', 'R': 'К', 'T': 'Е', 'Y': 'Н', 'U': 'Г', 'I': 'Ш', 'O': 'Щ', 'P': 'З',
	'{': 'Х', '}': 'Ъ', 'A': 'Ф', 'S': 'Ы', 'D': 'В', 'F': 'А', 'G': 'П', 'H': 'Р', 'J': 'О', 'K': 'Л', 'L': 'Д',
	':': 'Ж', '"': 'Э', 'Z': 'Я', 'X': 'Ч', 'C': 'С', 'V': 'М', 'B': 'И', 'N': 'Т', 'M': 'Ь',
	'<': 'Б', '>': 'Ю', '?': ',', '~': 'Ё',
	'@': '"', '#': '№', '$': ';', '^': ':', '&': '?',
}

var ruEn map[rune]rune

func init() {
	ruEn = make(map[rune]rune, len(enRu))
	for k, v := range enRu {
		ruEn[v] = k
	}
}

// Convert toggles each rune between Russian and Latin keyboard layouts.
func Convert(s string) string {
	runes := []rune(s)
	for i, ch := range runes {
		if mapped, ok := ruEn[ch]; ok {
			runes[i] = mapped
			continue
		}
		if mapped, ok := enRu[ch]; ok {
			runes[i] = mapped
		}
	}
	return string(runes)
}
