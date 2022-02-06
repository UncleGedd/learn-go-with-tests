package arrays

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// SumAll returns the sum of each slice passed in
func SumAll(numbersToSum ...[]int) (sum []int) {
	var sums []int
	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}
