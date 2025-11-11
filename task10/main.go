package main

import "fmt"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	var a, b int

	if _, err := fmt.Scan(&a, &b); err != nil {
		return
	}

	if b == 0 {
		fmt.Println("Devision by zero")
		return
	}

	q, r := divmod(a, b)
	fmt.Println(q, r)
}
