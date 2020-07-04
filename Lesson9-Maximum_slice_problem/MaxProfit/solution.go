package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func max(m, n int) (int, bool) {
	if m > n {
		return m, true
	}
	return n, false
}

func min(m, n int) (int, bool) {
	if m < n {
		return m, true
	}
	return n, false
}

func Solution(A []int) int {
	// write your code in Go 1.4
	result := 0
	minPrice := 0
	for i, v := range A {
		if i == 0 {
			minPrice = v
		} else {
			if p, ok := min(minPrice, v); !ok {
				minPrice = p
				continue
			}
			result, _ = max(result, v-minPrice)
		}
	}
	return result
}
