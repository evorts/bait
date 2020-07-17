package section3

import "fmt"

func SelectionSort(arr []int) {
	fmt.Println(arr)
	size := len(arr)
	var i, j, min int
	for i = 0; i < size-1; i++ {
		min = i
		for j = i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
}
