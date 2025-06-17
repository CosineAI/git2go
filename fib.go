// fib.go
// Package main provides a Fibonacci sequence generator using Go channels.
package main

import "fmt"

// FibGenerator returns a channel that produces Fibonacci numbers up to n values.
// If n is 0, it generates an infinite sequence (caller must break).
func FibGenerator(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 0, 1
		count := 0
		for {
			if n > 0 && count >= n {
				return
			}
			ch <- a
			a, b = b, a+b
			count++
		}
	}()
	return ch
}

// Example usage
func main() {
	fmt.Println("First 10 Fibonacci numbers:")
	for num := range FibGenerator(10) {
		fmt.Print(num, " ")
	}
	fmt.Println()
}