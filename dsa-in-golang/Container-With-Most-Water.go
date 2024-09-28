package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxarea := 0

	for left < right {
		ht := float64(min(height[left], height[right]))
		maxarea = int(math.Max(float64(maxarea), ht*float64(right-left)))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return maxarea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(height)
	fmt.Println(result)
}
