package main

import "fmt"

var n int

func fibonacci(n int) {
	a := 1
	b := 1
	c := b
	var i int = 1
	fmt.Println("n =", n)
	fmt.Printf("%d %d", a, b)
	for true {
		c = b
		b = a + b
		i += 1
		if i >= n {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}

func main() {
	fibonacci(10)
}
