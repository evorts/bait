package section3

func SumCollection(data []int) int {
	size, total := len(data), 0
	for idx := 0; idx < size; idx++ {
		total += data[idx]
	}
	return total
}