func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0

	for _, v := range prices {
		if v-min > profit {
			profit = v - min
		}

		if v < min {
			min = v
		}
	}

	return profit
}