package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func findPeaks(input []int, output []bool) {
	for i := 1; i < len(input)-1; i++ {
		if input[i] > input[i-1] && input[i] > input[i+1] {
			output[i] = true
		}
	}
}

func findNextPeaks(input []bool, output []int) {
	N := len(input)
	output[N-1] = -1 // There is nothing after last point.
	for i := N - 2; i >= 0; i-- {
		if input[i] {
			output[i] = i
		} else {
			output[i] = output[i+1]
		}
	}
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	peaks := make([]bool, N)
	nextPeak := make([]int, N)
	findPeaks(A, peaks)
	findNextPeaks(peaks, nextPeak)

	result := 0
	for i := 1; i*(i-1) <= N; i++ {
		p := 0
		n := 0
		for p < N && n < i {
			p = nextPeak[p]
			if p == -1 {
				break
			}
			n++
			p += i
		}
		result = max(result, n)
	}
	return result
}
