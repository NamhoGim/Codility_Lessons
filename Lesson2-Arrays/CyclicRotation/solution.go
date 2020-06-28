package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	// write your code in Go 1.4
	result := make([]int, len(A))
	for i, n := range A {
		result[(i+K)%len(A)] = n
	}
	return result
}
