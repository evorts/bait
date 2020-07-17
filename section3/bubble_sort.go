package section3

import "fmt"

func BubbleSort(arr []int, comp func(int, int) bool) {
	fmt.Println(arr)
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				/* Swapping */
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	fmt.Println(arr)
}

func BubbleSort2(arr []int, comp func(int, int) bool) {
	fmt.Println(arr)
	size := len(arr)
	swapped := 1
	for i := 0; i < (size-1) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = 1
			}
		}
	}
	fmt.Println(arr)
}
