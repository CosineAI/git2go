// Package fib provides a Fibonacci number generator.
package main

import "fmt"

// FibSlice returns a slice containing the first n Fibonacci numbers.
func FibSlice(n int) []int {
	if n <= 0 {
		return []int{}
	}
	fibs := make([]int, n)
	fibs[0] = 0
	if n > 1 {
		fibs[1] = 1
		for i := 2; i < n; i++ {
			fibs[i] = fibs[i-1] + fibs[i-2]
		}
	}
	return fibs
}

// FibGenerator returns a channel through which Fibonacci numbers are sent, up to count.
func FibGenerator(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 0, 1
		for i := 0; i < count; i++ {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}

// Example usage:
func main() {
	fmt.Println("First 10 Fibonacci numbers (slice):", FibSlice(10))

	fmt.Print("First 10 Fibonacci numbers (channel): ")
	for num := range FibGenerator(10) {
		fmt.Print(num, " ")
	}
	fmt.Println()
}