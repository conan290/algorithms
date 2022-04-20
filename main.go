package main

import "fmt"

func main() {
	var total int

	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Printf("The total is: %v", total)
}
