package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Solution(A []int) int {
	// write your code in Go 1.4
	result := 2147483647
	for i := 1; i < len(A); i++ {
		A[i] += A[i-1]
	}

	N := len(A) - 1
	for i := 1; i < len(A); i++ {
		result = min(result, abs(A[i-1]-(A[N]-A[i-1])))
	}
	return result
}
