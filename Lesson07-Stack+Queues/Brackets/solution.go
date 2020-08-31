package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	// write your code in Go 1.4
	s := []rune{}
	for _, r := range S {
		if r == '(' || r == '[' || r == '{' {
			s = append(s, r)
		} else {
			if r == ')' {
				if len(s) > 0 && s[len(s)-1] == '(' {
					s = s[:len(s)-1]
				} else {
					return 0
				}
			} else if r == ']' {
				if len(s) > 0 && s[len(s)-1] == '[' {
					s = s[:len(s)-1]
				} else {
					return 0
				}
			} else if r == '}' {
				if len(s) > 0 && s[len(s)-1] == '{' {
					s = s[:len(s)-1]
				} else {
					return 0
				}
			}
		}
	}
	if len(s) > 0 {
		return 0
	}
	return 1
}
