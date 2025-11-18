package main

import "fmt"

func main() {
	var n int
	var chet, znak string

	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	if n % 2 == 0 {
		chet = "even"
	} else {
		chet = "odd"
	}

	if n > 0 {
		znak = "positive"
	} else if n < 0 {
		znak = "negative"
	} else {
		fmt.Println("zero")
		return
	}

	fmt.Printf("%s %s\n", znak, chet)
}
