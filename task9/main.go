package main

import "fmt"

func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}

	return s
}

func main() {
	var s, num int

	for {
		fmt.Scan(&num)
		if num == 0 {break}
		s = sum(num, s)
	}
	
	fmt.Printf("%d\n", s)
}
