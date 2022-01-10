package leetcode

import (
	"fmt"
	"testing"
)

type para4 struct {
	num1 []int
	num2 []int
}

type ans4 struct {
	one float64
}

type question4 struct {
	para4
	ans4
}

func TestProblem4(t *testing.T) {
	qs := []question4 {
		{
			para4{[]int{1,3}, []int{2, 4}},
			ans4{2.5},
		},
		{
			para4{[]int{1, 2}, []int{3, 4}},
			ans4{2.5},
		},
		{
			para4{[]int{5}, []int{}},
			ans4{5},
		},
		{
			para4{[]int{}, []int{}},
			ans4{0},
		},
		{
			para4{[]int{1,3}, []int{2}},
			ans4{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 4------------------------\n")

	for _, q := range qs {
		result := findMedianSortedArrays(q.para4.num1, q.para4.num2)
		if q.one != result {
			fmt.Printf("Error:【input】:%v\t【output】:%v\t【Expected】：%v\n", q.para4, result, q.ans4)
		} else {
			fmt.Printf("OK:【input】:%v\t【output】:%v\t【Expected】：%v\n", q.para4, result, q.ans4)
		}
	}
	//fmt.Printf("\n\n\n")
}

