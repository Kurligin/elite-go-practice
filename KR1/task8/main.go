package main

import "fmt"

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	if n < 1 {
		fmt.Println("Must be not > 0")
	}

	var prev, curr int
	prev = 0
	curr = 1

	fmt.Printf("0 1 ")

	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", prev + curr)
		prev, curr = curr, prev + curr

	}
}
