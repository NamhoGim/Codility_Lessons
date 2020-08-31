package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(H []int) int {
	// write your code in Go 1.4
	result := 0
	s := []int{}
	for _, h := range H {
		for len(s) > 0 && s[len(s)-1] > h {
			s = s[:len(s)-1]
		}
		if len(s) == 0 || (len(s) > 0 && s[len(s)-1] < h) {
			result++
			s = append(s, h)
		}
	}
	return result
}
