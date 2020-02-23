package integers

import "fmt"

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	sum := x + y
	fmt.Println("Return value of Add: \n", sum)
	return sum
}
