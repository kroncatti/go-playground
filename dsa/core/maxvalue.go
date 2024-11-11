package dsaCore

func MaxValue(numbers []int) int {
	var maxVal int
	for _, item := range numbers {
		if maxVal < item {
			maxVal = item
		}
	}
	return maxVal
}
