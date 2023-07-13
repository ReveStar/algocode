package tpt

import "sort"

func Get3Sum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] <= nums[b]
	})
	result := make([][]int, 0)
	for k := 0; k < len(nums); k++ {
		i, j := k+1, len(nums)-1
		for i < j {
			if nums[i]+nums[j] == -nums[k] {
				re := []int{nums[i], nums[j], nums[k]}
				result = append(result, re)
				for j-1 >= 0 && nums[j] == nums[j-1] {
					j--
				}
				for i+1 < len(nums) && nums[i] == nums[i+1] {
					i++
				}
				j--
				i++
			} else if nums[i]+nums[j] > -nums[k] {
				j--
			} else {
				i++
			}

		}
		for k+1 < len(nums) && nums[k] == nums[k+1] {
			k++
		}
	}
	return result
}

func ValidTriangleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	if len(nums) < 3 {
		return 0
	}
	result := 0
	for k := 0; k < len(nums); k++ {
		i := k + 1
		j := len(nums) - 1
		for i < j {
			if nums[i]+nums[j] > nums[k] {
				result = result + (j - i)
				i++
			} else {
				j--
			}
		}
	}
	return result
}

func LongestSubstringWithoutRepeatString(s string) int {
	if len(s) == 0 {
		return 0
	}
	table := make(map[byte]int)
	i := 0
	j := 0
	result := 0
	for j < len(s) {
		index, ok := table[s[j]]
		if ok && index >= i {
			diff := j - i
			if diff > result {
				result = diff
			}
			i = index + 1
		}
		table[s[j]] = j
		j++
	}
	if j-i > result {
		result = j - i
	}
	return result
}
