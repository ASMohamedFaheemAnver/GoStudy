package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	n := 4096
	start := time.Now()
	result := fibonacci(n)
	elapsed := time.Since(start)
	fmt.Printf("Fibonacci(%d) = %d\n", n, result)
	fmt.Printf("Time taken: %s\n", elapsed)
}
