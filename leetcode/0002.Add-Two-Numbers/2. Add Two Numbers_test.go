package leetcode

import (
	"fmt"
	"testing"

	"github.com/derrick658/LeetCode/structures"
)

type para2 struct {
	one     []int
	another []int
}

type ans2 struct {
	one []int
}

type question2 struct {
	para2
	ans2
}

func Test_Problem2(t *testing.T) {
	qs := []question2{
		{
			para2{[]int{}, []int{}},
			ans2{[]int{}},
		},
		{
			para2{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans2{[]int{7, 0, 8}},
		},
		{
			para2{[]int{0}, []int{0}},
			ans2{[]int{0}},
		},
		{
			para2{[]int{9,9,9,9,9,9,9}, []int{9,9,9,9}},
			ans2{[]int{8,9,9,9,0,0,0,1}},
		},
		{
			para2{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans2{[]int{0, 0, 0, 0, 0, 1}},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")
	for _, q := range qs {
		_, p := q.ans2, q.para2
		fmt.Printf("【input】:%v       【output】:%v\n",
			p,
			structures.List2Ints(
				addTwoNumbers(structures.Ints2List(p.one), structures.Ints2List(p.another)),
			),
		)
	}
}
