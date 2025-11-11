package main

import "fmt"

func main() {
	var a, b int

	if _, err := fmt.Scan(&a, &b); err != nil {
		return
	}

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Printf("%.4f\n", float64(a) / float64(b))
	fmt.Println(a % b)
}
