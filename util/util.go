package util

func Larger(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 已排序列表，二分查找目标，返回索引，若未找到返回-1
func Search2Split(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			i = mid
		}
		if nums[mid] > target {
			j = mid
		}
	}
	return -1
}
