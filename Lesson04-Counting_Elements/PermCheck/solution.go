package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	maxVal := 0
	table := map[int]bool{}
	for _, n := range A {
		table[n] = true
		if maxVal < n {
			maxVal = n
		}
	}
	if maxVal == len(A) {
		for i := 1; i <= maxVal; i++ {
			if _, ok := table[i]; !ok {
				return 0
			}
		}
		return 1
	}
	return 0
}
