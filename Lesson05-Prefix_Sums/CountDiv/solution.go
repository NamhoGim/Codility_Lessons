package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
    // write your code in Go 1.4
    var a1, ak int
    if A % K == 0 {
        a1 = A
    } else {
        a1 = (A/K + 1)*K
    }
    ak = (B/K)*K
    
    return (ak - a1)/K + 1
}