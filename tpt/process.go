package tpt

import (
	"pcoding/excealgo/util"
	"sort"
)

func ContainerWithMostWater(heights []int) int {
	total := 0
	left := 0
	right := len(heights) - 1
	for left < right {
		s := util.Smaller(heights[left], heights[right])
		size := s * (right - left)
		if size > total {
			total = size
		}
		if heights[left] <= heights[right] {
			left++
		} else {
			right--
		}
	}
	return total
}

func ContainerWithMostWaterBadSolve(heights []int) int {
	total := 0
	for i := range heights {
		j := i + 1
		for ; j < len(heights); j++ {
			s := util.Smaller(heights[i], heights[j])
			total += (j - i) * s
		}
	}
	return total
}

func KDiffPairInAnArray(nums []int, k int) int {
	result := make(map[int]int)
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] <= nums[j]
	})
	i := 0
	j := 0

	for i <= j && j < len(nums) {
		if i == j {
			j++
			continue
		}
		s := nums[j] - nums[i]
		if s < k {
			j++
		} else if s == k {
			result[nums[i]] = nums[j]
			j++
		} else {
			i++
		}

	}
	return len(result)
}

func KDiffPairInAnArrayBad(nums []int, k int) int {
	result := make(map[int]int)
	for i := range nums {
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[i] <= nums[j] {
				if nums[j]-nums[i] == k {
					result[nums[i]] = nums[j]
				}
			} else {
				if nums[i]-nums[j] == k {
					result[nums[j]] = nums[i]
				}
			}
		}
	}
	return len(result)
}

func MinimumWindowSubstring(s string, t string) string {
	tcount := make(map[byte]int)
	for c := range t {
		_, ok := tcount[t[c]]
		if !ok {
			tcount[t[c]] = 1
		} else {
			tcount[t[c]]++
		}
	}
	rcount := make(map[byte]int)
	rlen := 0
	result := ""
	tlen := len(t)

	i, j := 0, 0
	for i <= j && j < len(s) {
		_, ok := tcount[s[j]]
		if ok {
			rcount[s[j]]++
			if rcount[s[j]] <= tcount[s[j]] {
				rlen++
			}

		} else {
			j++
			continue
		}
		for rlen == tlen {
			if result == "" {
				result = s[i : j+1]
			}
			if j-i+1 <= len(result) {
				result = s[i : j+1]
			}
			_, ok := tcount[s[i]]
			if ok {
				rcount[s[i]]--
				if rcount[s[i]] < tcount[s[i]] {
					rlen--
				}
			}
			i++
		}
		j++
	}
	return result
}
