package main

import (
	"fmt"
	"math"
)

// Find second max in a list

func GetSecondMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := -1 * math.MaxInt
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	secondMax := -1 * math.MaxInt

	for _, n := range nums {
		if n > secondMax && n < max {
			secondMax = n
		}
	}

	return secondMax
}

func main() {
	// expetcted 3
	list := []int{1, 2, 3, 4, -1, -2, -3}

	fmt.Println(GetSecondMax(list))
}
