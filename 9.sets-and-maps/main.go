package main

import "fmt"

// Example Problem - Set:
// Given two arrays, implement a function to find their intersection (elements that appear in both arrays).

func intersection(a, b []int) []int {
	res := []int{}
	set := make(map[int]bool)

	for _, n := range a {
		set[n] = true
	}

	for _, n := range b {
		if set[n] {
			res = append(res, n)
			delete(set, n)
		}
	}

	return res
}

// Example Problem - Map:
// Implement a function to count the frequency of each element in an array.

func countFreq(num []int) map[int]int {
	result := make(map[int]int)

	for _, n := range num {
		result[n]++
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	intersect := intersection(nums1, nums2)
	fmt.Println("Intersection:", intersect) // [2]

	nums := []int{1, 2, 3, 1, 2, 1, 3, 4, 5}

	frequency := countFreq(nums)
	fmt.Println("Frequency Map:", frequency)
}
