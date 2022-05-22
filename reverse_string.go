package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	// solution goes here

	for _, name := range input {
		output = fmt.Sprint(string(name)) + output
	}
	return output
}
