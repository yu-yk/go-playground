package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// define left right boundary
	left, right := 0, len(nums)-1

	// select random pivot
	pivotIndex := rand.Intn(len(nums))

	// swap the pivot with the right pointer
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	// Pile elements smaller than the pivot on the left
	// left pointer is always the position after last smaller elem
	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	nums[left], nums[right] = nums[right], nums[left]

	// Go down the rabbit hole
	quickSort(nums[:left])
	quickSort(nums[left+1:])

	return nums
}

func main() {
	n := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(n)
	fmt.Println(n)

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)
}
