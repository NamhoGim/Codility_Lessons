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

func Solution(N int) int {
    // write your code in Go 1.4
    result, count := 0, 0
    flag := false
    for i := 31; i >= 0; i-- {
        if !flag && (N & (1 << uint(i)) != 0) {
            flag = true
        } else if flag && (N & (1 << uint(i)) == 0) {
            count++
        } else if flag && (N & (1 << uint(i)) != 0) {
            result = max(result, count)
            count = 0
        }
    }
    return result
}
