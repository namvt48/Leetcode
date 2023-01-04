func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		v, oke := hm[target-nums[i]]

		if oke && i != v {

			return []int{i, v}
		}

		hm[nums[i]] = i
	}

	return nil
}
