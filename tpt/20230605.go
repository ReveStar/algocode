package tpt

func Get4Sum(nums []int, target int) [][]int {
	/*
	 * 1. 四元组无顺序限制
	 * 2. 思路类似3sum但是由于外层两层循环，需要尽可能剪枝降低复杂度
	 * 3. 先排序
	 * 4. 外层两层循环，内部双指针遍历获取结果
	 * 5. 剪枝排除重复的选项
	 */
	result := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return nil
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k, l := j+1, n-1
			atarget := target - nums[i] - nums[j]
			for k < l {
				sum := nums[k] + nums[l]
				if sum == atarget {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					for k+1 < n && nums[k] == nums[k+1] {
						k++
					}
					for l-1 >= 0 && nums[l] == nums[l-1] {
						l--
					}
					k++
					l--
				} else if sum < atarget {
					k++
				} else {
					l--
				}
			}
			for j+1 < n && nums[j] == nums[j+1] {
				j++
			}
		}
		for i+1 < n && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}

func MaxConsecutiveOnesIII(nums []int, k int) int {
	/*
	 * 1. 右指针移动至无法移动，即k被消耗完
	 * 2. 比较大小并记录
	 * 3. 左指针移动至k 大于0
	 * 4. 继续移动右指针
	 * 5. 循环至结束
	 */
	result := 0
	n := len(nums)
	count := k
	i := 0
	for j := 0; j < n; j++ {
		if nums[j] == 0 {
			for count <= 0 && i <= j {
				if nums[i] == 0 {
					count++
				}
				i++
			}
			if count > 0 {
				count--
				if j-i+1 > result {
					result = j - i + 1
				}
			}
		} else {
			if j-i+1 > result {
				result = j - i + 1
			}
		}
	}
	return result
}

func LongestSubstringWithAtMostKDistinctCharacters(s string, k int) string {
	/*
	 * 1. 记录子串中每个字符个数
	 * 2. 记录子串中出现的不同字符的数量
	 * 3. 移动右指针，至不同字符数量大于K
	 * 5. 移动左指针并减少对应字符的数量,至到不同字符的数量小于k
	 * 6. 继续移动右指针,直结束，并比较循环结束后最后一个子串和已有最大子串的大小
	 */
	result := ""
	index := make([]int, 256)
	n := len(s)
	count := 0
	i := 0
	j := 0
	for ; j < n; j++ {
		index[s[j]]++
		if index[s[j]] == 0 {
			count++
			if count > k {
				if len(result) < j-i {
					result = s[i:j]
				}
				for count > k {
					index[s[i]]--
					if index[s[i]] == 0 {
						count--
					}
					i++
				}
			}
		}
	}
	if j-i > len(result) {
		result = s[i:j]
	}
	return result
}
