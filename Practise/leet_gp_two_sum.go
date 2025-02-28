package Practise

import "fmt"

var count int = 0

func main() {

	Nums := []int{10, 2, 3, 4, 5}
	Target := 5

	fmt.Println(twoSum(Nums, Target), "Итераций: ", count)
}
func twoSum(nums []int, target int) []int {
	var Output []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			count++
			if nums[i]+nums[j] == target {
				Output = []int{nums[i], nums[j]}
				return Output
			}
		}
	}
	return nil
}
