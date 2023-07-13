package tpt

func Get3SumSmaller(nums []int, target int) int {
	/*
	 * 1. 排序数组
	 * 2. 外层循环i
	 * 3. 内层寻找小于target的可能值
	 * 4. 若j + k 小于target,则[j,k]之间都是小于target的值
	 * 5. 记录符合要求的个数
	 */
	result := 0
	return result
}

func GrumpyBookstoreOwner(customers []int, grumpy []int, minutes int) int {
	/*
	 * 1. 右指针遍历grumpy至minutes为0,并记录由0变1增加的数量,比较此滑动窗口新增数量是够更大.
	 * 2. 左指针右移至minutes大于0,并减去由1变0减少的数量
	 * 3. 继续右移右指针,重复1,2至右指针遍历完成
	 * 4. 遍历结束后比较新增数量是不是更大.输出最大值
	 */
	result := 0
	return result
}

func SubarraysWithKDifferentInteger(nums []int, k int) []int {
	/*
	 * 1.at most K - at most K-1 得到K不同整数的子串数量
	 * 2. 需要实现at most K 的函数
	 */
	result := make([]int, 0)
	return result
}

func SubarraysWithKDifferentIntegerFind(nums []int, k int) []int {
	/*
	 * 1. 右指针右移到K个不同整数,计数加一,依次右移右指针直大于K个不同整数,并计数
	 * 2. 右移左指针直K个不同整数,计数加一,依次右移右指针至大于K个不同整数,并计数,循环步骤2
	 * 3. 至遍历结束,返回结果
	 */
	result := make([]int, 0)
	return result
}
