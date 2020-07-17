package section2

import "fmt"

func Array() {
	var arr [10]int
	fmt.Println(arr)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	fmt.Printf("Length of array: %d", len(arr))
}
