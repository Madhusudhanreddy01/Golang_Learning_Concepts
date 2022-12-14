package main

import "fmt"

/*
https://oj.leetcode.com/problems/two-sum/
Given an array of integers, find two numbers such that they add up to a specific target number.
The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution.
Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2})
*/
func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println((twoSum(nums, 9)))
}

// func twoSum(nums []int, target int) (index1, index2 int) {
// 	for i1, _ := range nums {
// 		for i2 := i1 + 1; i2 < len(nums); i2++ {
// 			if nums[i1]+nums[i2] == target {
// 				index1 = i1 + 1
// 				index2 = i2 + 1
// 			}
// 		}
// 	}
// 	return index1, index2
// }

func twoSum(nums []int, target int) []int {
	var a []int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums)-1; j++ {
			if nums[i]+nums[j+1] == target {
				a = append(a, i, j+1)
			}
		}
	}
	return a
}
