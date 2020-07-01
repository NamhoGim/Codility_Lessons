package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	// write your code in Go 1.4
	flag := false
	level := 0
	for _, r := range S {
		if !flag && r == '(' {
			flag = true
			level++
		} else if !flag && r == ')' {
			return 0
		} else if flag && r == '(' {
			level++
		} else if flag && r == ')' {
			level--
			if level == 0 {
				flag = false
			}
		}
	}
	if level != 0 {
		return 0
	}
	return 1
}
