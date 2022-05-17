package reverse_string

func ReverseString(input string) (output string) {
	var runes []rune = []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		output += string(runes[i])
	}

	return output
}
