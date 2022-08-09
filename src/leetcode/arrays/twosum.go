//two sum problem
//[1,7,8,15] target=9 ... so 1+8 = 9 so output is [0,2]

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(append(res, i), j)
				return res
				break
			}
		}
	}
	return res
}

func main() {

	temp := []int{1, 2, 3, 6}
	target := 5
	res := twoSum(temp, target)
	fmt.Println(res)
}
