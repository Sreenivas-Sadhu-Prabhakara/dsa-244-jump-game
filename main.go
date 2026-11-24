package main

import "fmt"

// CanJump reports whether the last index is reachable from the first, where
// nums[i] is the maximum hop from i. One pass tracking the farthest
// reachable index; falling behind it means being stranded.
func CanJump(nums []int) bool {
	reach := 0
	for i, v := range nums {
		if i > reach {
			return false // this index is beyond anything reachable
		}
		if i+v > reach {
			reach = i + v
		}
	}
	return true
}

func main() {
	fmt.Println(CanJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(CanJump([]int{3, 2, 1, 0, 4}))
}
