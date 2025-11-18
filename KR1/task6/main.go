package main

import "fmt"

func main() {
	var score int
	
	if _, err := fmt.Scan(&score); err != nil {
		return
	}

	switch {
	case score < 0 || score > 100:
		fmt.Println("invalid")
	case score == 100:
		fmt.Println("A")
	default:
		switch score / 10 {
		case 9:
			fmt.Println("A")
		case 8:
			fmt.Println("B")
		case 7:
			fmt.Println("C")
		case 6:
			fmt.Println("D")
		default:
			fmt.Println("F")
		}
	}
}