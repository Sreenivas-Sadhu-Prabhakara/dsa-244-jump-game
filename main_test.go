package main

import "testing"

func TestCanJump(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true}, // already standing on the last index
		{[]int{1, 0}, true},
		{[]int{0, 1}, false},
		{[]int{5, 0, 0, 0}, true},
		{[]int{1, 1, 1, 0, 1}, false},
	}
	for _, c := range cases {
		if got := CanJump(c.in); got != c.want {
			t.Errorf("CanJump(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
