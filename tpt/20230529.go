package tpt

import "sort"

func ClosetTarget3Sum(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})
	diff := 600000
	result := 0
	for k, v := range nums {
		atarget := target - v
		i := k + 1
		j := len(nums) - 1
		for i < j {
			s := nums[i] + nums[j]
			if s == atarget {
				return target
			}
			ad := absDiffInt(s, atarget)
			if ad < diff {
				result = s + v
				diff = ad
			}
			if s < atarget {
				i++
			} else {
				j--
			}
		}
	}
	return result
}

func BinarySubarraysWithSumSlidWindow(binarys []int, goal int) int {
	if len(binarys) == 0 {
		return 0
	}
	count := 0
	aftZoreNum := make([]int, len(binarys))
	for i := len(binarys) - 1; i >= 0; i-- {
		aftZoreNum[i] = count
		if binarys[i] == 0 {
			count++
		} else {
			count = 0
		}
	}
	result := 0
	i, j := 0, 0
	sum := 0
	for i < len(binarys) {
		for j <= i || (j < len(binarys) && sum < goal) {
			sum += binarys[j]
			j++
		}
		if sum == goal {
			result = result + aftZoreNum[j-1] + 1
		}
		sum -= binarys[i]
		i++
	}
	return result
}

func BinarySubarraysWithSumPrefix(binarys []int, goal int) int {
	if len(binarys) == 0 {
		return 0
	}
	result := 0
	index := make(map[int]int, 0)
	index[0] = 1
	table := make([]int, len(binarys))
	for i, v := range binarys {
		if i == 0 {
			table[i] = v
		} else {
			table[i] = table[i-1] + v
		}
		index[table[i]]++
		result = result + index[(table[i]-goal)]
	}
	return result
}

func LongestSubstringWithAtMostTwoDistinctCharacters(s string) string {
	result := ""
	if s == "" {
		return ""
	}
	loc := make(map[byte]int)
	pre := s[0]
	sec := s[0]
	loc[pre] = 0
	loc[sec] = 0
	for j := 0; j < len(s); j++ {
		if s[j] == pre {
			continue
		} else {
			sec = s[j]
			loc[sec] = j
			break
		}
	}
	if pre == sec {
		return s
	}
	i := 0
	j := loc[sec]
	for j < len(s) {
		if s[j] != pre && s[j] != sec {
			if j-i > len(result) {
				result = s[i:j]
			}
			if loc[pre] < loc[sec] {
				i = loc[pre] + 1
				pre = s[j]
				loc[pre] = j
			} else {
				i = loc[sec] + 1
				sec = s[j]
				loc[sec] = j
			}
		} else if s[j] == pre {
			loc[pre] = j
		} else {
			loc[sec] = j
		}
		j++
	}
	if j-i > len(result) {
		result = s[i:j]
	}
	return result
}
