package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "ababcdabc"

	if len(s) == 0 {
		fmt.Println(0)
		return
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
			fmt.Println(left, right)
		} else {
			bitSet[s[right]] = true
			right++
		}

		result = max(result, right-left)
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	fmt.Printf("result = %d\n", result)
	return
}

/*
	分析：
		第一次循环结束后：indexes:{a: 0}; 					left:1; right:0; res: 1
		第二次循环结束后：indexes:{a: 0, b: 1}; 			left:2; right:0; res: 2
		第三次循环结束后：indexes:{a: 2, b: 1};			left:3; right:1; res: 2
		第四次循环结束后：indexes:{a: 2, b: 3}; 			left:4; right:2; res: 2
		第五次循环结束后：indexes:{a: 2, b: 3, c:4}; 		left:5; right:2; res: 3
		第六次循环结束后：indexes:{a: 2, b: 3, c:4, d:5}; 	left:6; right:2; res: 4
		第七次循环结束后：indexes:{a: 6, b: 3, c:4, d:5}; 	left:7; right:3; res: 4
		第八次循环结束后：indexes:{a: 6, b: 7, c:4, d:5}; 	left:8; right:4; res: 4
		第九次循环结束后：indexes:{a: 6, b: 7, c:8, d:5}; 	left:9; right:5; res: 4

		结论：这相当于求解哪段距离最大，中间不能有重复字符。将所有字符遍历一下，记录每个字符最新
	的位置，若果有重复，更新位置。每次循环一次，记录一下距离有多大，并且和上次的距离进行比较
*/
