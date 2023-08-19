package array

func HasArrayItem(item int, items []int) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
