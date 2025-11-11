package main

import "fmt"

func main() {
	var PI float64 = 3.14159
	var r float64
	
	fmt.Scan(&r)
	
	area := PI * r * r
	fmt.Printf("%.3f\n", area)
}