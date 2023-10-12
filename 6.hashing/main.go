package main

import "fmt"

// Problem Statement:
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

func twoSum(num []int, sum int) (int, int) {
	seen := map[int]int{}

	for i, n := range num {
		want := sum - n

		second, ok := seen[want]
		if ok {
			return i, second
		}

		seen[n] = i
	}

	return -1, -1
}

func main() {
	nums := []int{4, 3, 2, 7, 11, 15}
	target := 9

	i, j := twoSum(nums, target)
	fmt.Printf("Indices of the two numbers: %d %d\n", i, j)
}
