package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int) int {
	// write your code in Go 1.4
	var result, maxVal int
	exist := map[int]bool{}
	for _, n := range A {
		if n < 0 {
			continue
		}
		if _, ok := exist[n]; !ok {
			exist[n] = true
			maxVal = max(maxVal, n)
		}
	}

	for i := 1; i <= maxVal; i++ {
		if _, ok := exist[i]; !ok {
			result = i
			break
		}
	}

	if result == 0 {
		return maxVal + 1
	}
	return result
}
