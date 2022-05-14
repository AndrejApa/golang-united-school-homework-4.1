package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	x := []rune(input)
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	output = string(x)
	return output
}
