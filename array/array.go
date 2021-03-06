package array

//给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
//你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用
//
//给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
func twoSum(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	var res []int
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				res = []int{i, j}
			}
		}
	}
	return res
}

func twoSumTwoMap(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	res := make([]int, 2)
	for i, v := range nums {
		w := target - v
		if j, ok := m[w]; ok {
			res[0] = j
			res[1] = i
			return res
		}
		m[v] = i
	}
	return res
}

//maxProfit 股票
//
// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	var profit int
	first, second := prices[0], prices[0]
	for i := 1; i < length; i++ {
		if second >= prices[i] {
			if second > first {
				profit = profit + (second - first)
			}
			first, second = prices[i], prices[i]
		} else if second < prices[i] {
			if i == length-1 {
				profit = profit + (prices[i] - first)
			}
			second = prices[i]
		}
	}
	return profit
}

//maxProfitPerfec 优雅的写法
func maxProfitPerfect(prices []int) int {
	var profit int
	for i := 1; i < len(prices); i++ {
		if num := prices[i] - prices[i-1]; num > 0 {
			profit += num
		}
	}
	return profit
}

//rotate 旋转数组
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	n := len(nums) - k%len(nums)
	tmp := append(nums[n:], nums[:n]...)
	for i, v := range tmp {
		nums[i] = v
	}
}

func rotate1(nums []int, k int) {
	if k == 0 {
		return
	}
	length := len(nums)
	n := k % length
	tmp := make([]int, length)
	copy(tmp, nums)
	for i := 0; i < length; i++ {
		if i+n-length >= 0 {
			nums[i+n-length] = tmp[i]
		} else {
			nums[i+n] = tmp[i]
		}
	}
}

//containsDuplicate
//给定一个整数数组，判断是否存在重复元素。
// 如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false
func containsDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = 1
	}
	return false
}

//singleNumber
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	m := make(map[int]int, length/2)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

//利用异或的特性 A^B^C^B^A = B^B^A^A^C = C
func singleNumberPerfect(nums []int) int {
	res := 0
	for _, n := range nums {
		res = res ^ n
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	data := []int{}
	m1 := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		if _, ok := m1[v]; ok && m1[v] > 0 {
			data = append(data, v)
			m1[v]--
		}
	}
	return data
}

//plusOne 给定一个非负整数组成的非空数组，在该数的基础上加一，返回一个新的数组。
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			return digits
		}
	}
	//第一位为0 特殊处理
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

//moveZeroes 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] == 0 {
			for j := i + 1; j < length; j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

func moveZeroes_1(nums []int) {
	length := len(nums)
	for i, k := 0, 0; i < length; i++ {
		if nums[i] != 0 {
			if i != k {
				nums[i], nums[k] = nums[k], nums[i]
			}
			k++
		}
	}
}

//判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
func isValidSudoku(board [][]byte) bool {
	dataC := [9][9]bool{}
	dataR := [9][9]bool{}
	dataS := [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			c := board[i][j] - '0' - 1
			if dataC[i][c] || dataR[c][j] || dataS[3*(i/3)+j/3][c] {
				return false
			}
			dataC[i][c] = true
			dataR[c][j] = true
			dataS[3*(i/3)+j/3][c] = true
		}
	}
	return true
}
