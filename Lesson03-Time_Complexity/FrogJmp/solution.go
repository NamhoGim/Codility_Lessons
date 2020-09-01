package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	if Y-X == 0 {
		return 0
	} else if (Y-X)%D == 0 {
		return (Y - X) / D
	}
	return (Y-X)/D + 1
}
