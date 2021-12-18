package leetcode

import (
	"fmt"
	"testing"
)

type para1 struct {
	nums   []int
	target int
}

type ans1 struct {
	one []int
}

type question1 struct {
	para1
	ans1
}

func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			para1{nums: []int{2, 7, 11, 15}, target: 9},
			ans1{[]int{0, 1}},
		},
		{
			para1{nums: []int{3, 2, 4}, target: 7},
			ans1{[]int{1, 2}},
		},
		{
			para1{nums: []int{3, 3}, target: 6},
			ans1{[]int{0, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v       【output】:%v\n", p, twoSum(p.nums, p.target))
	}
	fmt.Println()
}
