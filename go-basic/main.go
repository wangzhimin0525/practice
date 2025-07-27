package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums [5]int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return -1
}

func checkHuiWen(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}
	return x == reversedHalf || x == reversedHalf/10
}

func validBrackets(s string) bool {
	var stack []rune
	bracketMap := map[rune]rune{ // 闭括号到开括号的映射
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		if closing, ok := bracketMap[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		prefix = prefix[:j]
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func plusOne(digits []int) []int {
	n := len(digits)

	// 从最低位（数组末尾）开始遍历
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++ // 当前位 <9，直接 +1 并返回
			return digits
		}
		digits[i] = 0 // 当前位 =9，进位后变为 0
	}

	// 如果所有位都是 9，需要在最前面插入 1
	return append([]int{1}, digits...)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0 // 慢指针，指向当前不重复部分的最后一个元素
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j] // 将不重复元素放到前面
		}
	}
	return i + 1 // 新长度 = 最后一个不重复元素的索引 + 1
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 1. 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化结果数组
	merged := [][]int{intervals[0]}

	// 3. 遍历并合并
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		curr := intervals[i]

		if curr[0] > last[1] {
			// 当前区间与最后一个区间不重叠，直接加入
			merged = append(merged, curr)
		} else {
			// 当前区间与最后一个区间重叠，合并它们
			last[1] = max(last[1], curr[1])
		}
	}
	return merged
}

func twoSum(nums []int, target int) []int {
	// 创建一个哈希表，存储数值 -> 下标
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, ok := numMap[complement]; ok {
			// 如果 complement 存在于哈希表中，返回它们的下标
			return []int{idx, i}
		}
		// 否则，存入当前数值和下标
		numMap[num] = i
	}

	return nil
}

func main() {
	//nums := [5]int{1, 2, 2, 3, 3}
	//fmt.Println(singleNumber(nums))

	//fmt.Println(checkHuiWen(12321))

	//fmt.Println(validBrackets("()[]{}"))

	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	//fmt.Println(plusOne([]int{9, 3, 2, 9}))

	//nums1 := []int{1, 2, 2, 3, 3, 3}
	//length1 := removeDuplicates(nums1)
	//fmt.Println(length1, nums1[:length1])

	//intervals := [][]int{{1, 4}, {2, 3}}
	//fmt.Println(merge(intervals))

	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
