package leetcode

// 解法一 位图
// TODO: 没搞懂
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for left < len(s) {
		if bitSet[s[left]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}

		result = max(result, right-left)
		if left+result > len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法二 滑动窗口
// TODO: 没搞懂
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	var freq [256]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

// 解法三 滑动窗口-哈希桶，遍历输入字符串，每个字符串求解一下距离，并和之前的最大距离比较，大则更新，小不更新。
func lengthOfLongestSubstring2(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
