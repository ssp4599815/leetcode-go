package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

思路：
从第一个元素开始 循环的和数组中剩余的元素相加，
每循环一次结束后，从第二个元素开始，重复上面的过程
*/


func twoSum(nums []int, target int) []int {
	i := 0
	// 列表有多长就循环多少次
	for i < len(nums)-1 {
		// 获取 第一个元素，来和列表中所有的元素 相加j
		first := nums[i]
		for j := i + 1; j < len(nums); j++ {
			ret := first + nums[j]
			if ret == target {
				return []int{i, j}
			}
		}
		// 每次遍历完都会+1，从而让后面的列表少一个
		i++
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15, 16, 19, 22, 40, 90}
	target := 31
	fmt.Println(twoSum(nums, target))
}
