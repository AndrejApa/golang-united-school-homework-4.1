package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	for _, v := range input {
		output += string(v)
	}
	return output
}
