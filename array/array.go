package array

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
