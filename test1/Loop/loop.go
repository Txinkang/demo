package Loop

func Loop(nums []int, step int) {
	l := len(nums)
	for i := 0; i < step; i++ {
		for j := i; j < l; j += step {
			nums[j] = 4 //访问内存，并写入值
		}
	}
}
