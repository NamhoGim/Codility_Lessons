package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func Solution(A []int) int {
	// write your code in Go 1.4
	result := -2147483648
	tempSum := 0
	for i := range A {
		tempSum = max(A[i], tempSum+A[i])
		result = max(result, tempSum)
	}

	return result
}
