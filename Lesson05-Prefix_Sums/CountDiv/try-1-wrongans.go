package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	result := 0
	for i := K; i <= B; i = i + K {
		if i >= A && i <= B {
			result++
		}
	}
	return result
}
