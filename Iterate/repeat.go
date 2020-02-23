package iteration

import "fmt"

const repeatCount = 3

// Repeat returns a repeated character `n` number of times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	fmt.Println("Return value of Repeat: \n", repeated)
	return repeated
}
