### 1. [Two Sum](https://leetcode.com/problems/two-sum/)

Given an array of integers `nums` and an integer `target`, return *indices of the two numbers such that they add up to `target`*.

You may assume that each input would have **exactly one solution**, and you may not use the *same* element twice.

You can return the answer in any order.



**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

 

**Constraints:**

- `2 <= nums.length <= 104`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
- **Only one valid answer exists.**

 

**Follow-up:** Can you come up with an algorithm that is less than `O(n2) `time complexity?



### 题目大意

在数组中找出两个元素之和等于给定值的数字，结果返回两个元素在数组中的下表。



### 解题思路

这道题的最优的做法是时间复杂度是 O(n)。

顺序扫描数组，对每一个元素，在 map 中找到给定值的另一半数字，如果找到了，就直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫描到“另一半”数字的时候，再取出来返回结果。

