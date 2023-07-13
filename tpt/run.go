package tpt

import "fmt"

func minisubstring() {
	s := "ADOBECODEBANC"
	t := "ABC"
	r := MinimumWindowSubstring(s, t)
	fmt.Println(r)
}

func get3sum() {
	nums := []int{
		-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0,
	}
	r := Get3Sum(nums)
	fmt.Println(r)
}

func ValidateTriangle() {
	nums := []int{
		2, 2, 3, 4,
	}
	r := ValidTriangleNumber(nums)
	fmt.Println(r)
}

func closetTarget() {
	nums := []int{
		-1, 2, 1, -4,
	}
	target := 1
	r := ClosetTarget3Sum(nums, target)
	fmt.Println(r)
}

func binaryTarget() {
	nums := []int{
		1, 0, 1, 0, 1,
	}
	r := BinarySubarraysWithSumPrefix(nums, 2)
	fmt.Println(r)
}

func RunCase() {
	//	minisubstring()
	// get3sum()
	// ValidateTriangle()
	// closetTarget()
	binaryTarget()
}
