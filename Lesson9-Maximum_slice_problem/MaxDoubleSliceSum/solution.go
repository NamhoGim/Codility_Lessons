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
	N := len(A)
	sumL := make([]int, N)
	sumR := make([]int, N)
	for i := 1; i < N-1; i++ {
		sumL[i] = max(sumL[i-1]+A[i], 0)
	}
	for i := N - 2; i > 0; i-- {
		sumR[i] = max(sumR[i+1]+A[i], 0)
	}

	var result int
	for i := 1; i < N-1; i++ {
		result = max(result, sumL[i-1]+sumR[i+1])
	}
	return result
}
