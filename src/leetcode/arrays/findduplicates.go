// find duplicates for eg. [1,2,2,3,4,5,5]
// output [2,5] since only 2 and 5 are repeating more than once
package main

import "fmt"

func findDuplicate(num []int) []int {
	var res []int
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] == num[j] {
				res = append(res, num[i])
				break
			}
		}
	}
	return res
}

func main() {
	n := []int{4, 3, 2, 7, 8, 2, 3, 1}
	res := findDuplicate(n)
	fmt.Println(res)
}
