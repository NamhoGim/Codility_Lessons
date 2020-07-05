package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	peaks := []int{}
	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	for size := len(peaks); size > 0; size-- {
		if len(A)%size == 0 {
			blockSize := len(A) / size
			foundCnt := 0
			found := make([]bool, size)
			for _, peak := range peaks {
				if !found[peak/blockSize] {
					found[peak/blockSize] = true
					foundCnt++
				}
			}
			if foundCnt == size {
				return size
			}
		}
	}
	return 0
}
