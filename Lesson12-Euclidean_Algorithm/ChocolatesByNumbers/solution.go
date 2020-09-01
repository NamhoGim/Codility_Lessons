package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func Solution(N int, M int) int {
	lcm := N * M / gcd(N, M)
	return lcm / M
}
