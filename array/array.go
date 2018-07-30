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
