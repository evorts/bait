package section2

import "fmt"

func Slice() {
	var s []int
	for i := 1; i <= 17; i++ {
		s = append(s, i)
		fmt.Printf("%v :: len = %d, capacity = %d \n", s, len(s), cap(s))
	}
}
