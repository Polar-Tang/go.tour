package main

func sum(nums ...int) int {
	var acc2 int
	for i := 0; i < len(nums); i++ {
		acc2 = acc2 + nums[i]
	}
	return acc2
}
