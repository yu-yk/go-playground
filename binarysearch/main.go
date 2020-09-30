package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, 5, 7, 11}
	min := smallestDivisor(nums, 11)

	fmt.Println(min)

	// fmt.Println(feasible(3, 11, nums))
	// fmt.Println(2 / 3)
}

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, max(nums)

	for left <= right {
		mid := left + (right-left)/2

		if feasible(mid, threshold, nums) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func feasible(divisor, threshold int, nums []int) bool {
	sum := 0

	for _, n := range nums {
		sum += int(math.Ceil(float64(n) / float64(divisor)))
		if sum > threshold {
			return false
		}
	}
	// 4 : 1 1 2 3
	// 5 : 1 1 1 2
	return true
}

// https://leetcode.com/discuss/general-discussion/786126/python-powerful-ultimate-binary-search-template-solved-many-problems
// 收細search space, 直到收到left / left-1 係starting point meet condition
func binSearch(a []int, value int) bool {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2

		// condition
		if condition(a[mid], value) {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	if a[left] == value {
		return true
	}
	return false
}

func condition(n, value int) bool {
	if n >= value {
		return true
	}
	return false
}

// template
func max(nums []int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func min(nums []int) int {
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
