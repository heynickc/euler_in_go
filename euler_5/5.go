package main

import (
	"fmt"
)

func main() {
	lowNums := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for i, num := range lowNums {
		fmt.Printf("%v index; %v num", i, num)
	}
}
