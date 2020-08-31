package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	result := 0
	count := 0
	for _, n := range A {
		if n == 0 {
			count++
		} else {
			result += count
		}
	}
	if result > 1000000000 {
		return -1
	}
	return result
}
